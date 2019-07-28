package hkr;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.Scanner;

import org.bukkit.Material;
import org.bukkit.entity.Player;
import org.bukkit.inventory.ItemStack;
import org.bukkit.plugin.java.JavaPlugin;
import org.bukkit.potion.PotionEffect;
import org.bukkit.potion.PotionEffectType;

import hkr.ArmorEquipEventFiles.ArmorListener;
import hkr.ArmorEquipEventFiles.DispenserArmorListener;

public class ArmorSetBonusMain extends JavaPlugin
{
    private ArrayList<ArmorSet> armorSets = new ArrayList<>();

    private HashMap<Player,ArmorSet> activeBonus = new HashMap<>();



    @Override
    public void onEnable(){
        try {
            checkFiles();
        } catch (IOException e) {
            getLogger().info("Error checking files");
            e.printStackTrace();
        }
        try {
            loadConfig();

        } catch (FileNotFoundException e) {
            getLogger().info("Error loading config. Invalid perhaps?");
            e.printStackTrace();
        }

        new CommandInit(this).InitCommands();

        
        getServer().getPluginManager().registerEvents(new ArmorListener(getConfig().getStringList("blocked")), this);
        getServer().getPluginManager().registerEvents(new ArmorEquipListener(this), this);
        try {
            // Better way to check for this? Only in 1.13.1+?
            Class.forName("org.bukkit.event.block.BlockDispenseArmorEvent");
            getServer().getPluginManager().registerEvents(new DispenserArmorListener(), this);
        } catch (Exception ignored) {}

        

    }
    public boolean flush(){
        armorSets.clear();
        try {
            checkFiles();
            loadConfig();
        } catch (Exception e) {
            getLogger().warning("Invalid config file.");
            return false;
        }
        updateAll();
        return true;
    }

    private void updateAll() {
        for (Player player : getServer().getOnlinePlayers()) {
            checkForBonus(player);
        }
    }

    private void loadConfig() throws FileNotFoundException {
        getLogger().info("Loading config");
        Scanner sc = new Scanner(new File(this.getDataFolder() + "/armorSets.txt"));
        String name = "";
        Material[] armor = new Material[4];
        ArrayList<PotionEffect> effect = new ArrayList<>();
        
        boolean newSet = false;
        String line = sc.nextLine();
        while (sc.hasNextLine() || newSet) {
            // getLogger().info("I'm here: 1");
            if (newSet) {
                // getLogger().info("I'm here: newset");
                newSet = false;
                if (name != "") {
                    ArmorSet newArmorSet = new ArmorSet(name, armor, effect);
                    armorSets.add(newArmorSet);
                    name = "";
                    armor = new Material[4];
                    effect = new ArrayList<>();
                }
            } else {
                if (line.startsWith("#")) {
                    line = sc.nextLine();
                } else {
                    if (line.contains("SET")) {
                        // getLogger().info("I'm here: 2");
                        name = line.substring(0, line.length()-1);
                        line = sc.nextLine();
                    } 
                    if (line.contains("Armor")) {
                        // getLogger().info("I'm here: 3");
                        for (int y = 0; y < 4; y++) {
                            // getLogger().info("I'm here: 3."+y);
                            line = sc.nextLine();
                            armor[y] = Material.getMaterial(line.trim(), false);
                            if (armor[y] == null) {
                                armor[y] = Material.AIR;
                            }
                            
                        }
                        line = sc.nextLine();
                    } 
                    if (line.contains("Effects")) {
                        line = sc.nextLine();
                        while (!line.contains("_SET")) {
                            // getLogger().info("I'm here: 4");
                            String[] temp = line.split(",");
                            PotionEffectType pE = PotionEffectType.getByName(temp[0].trim());
                            Integer amp = 1;
                            try {
                                amp = Integer.parseInt(temp[1].trim());
                            } catch (NumberFormatException e) {
                                amp = 1;
                            }
                            if (pE != null) {
                            PotionEffect nP = new PotionEffect(pE, Integer.MAX_VALUE, amp);
                            effect.add(nP);
                            }
                            if (sc.hasNextLine()) {
                                line = sc.nextLine();
                            } else {
                                break;
                            }
                        }
                        // Marks the end of reading for that spesific set
                        // Makes sure it does not skip one line ahead
                        newSet = true;
                    }
                }
            }
        }
        sc.close();
    }

    @Override
    public void onDisable() {
        for (Player player : activeBonus.keySet()) {
            removeBonus(player);
        }
    }

    private void checkFiles() throws IOException {
        saveDefaultConfig();
        File f = new File(this.getDataFolder() + "/");
        if (!f.exists()) {
            f.mkdir();
        }
        File fc = new File(this.getDataFolder() + "/armorSets.txt");
        if (!fc.exists()){
            saveResource("armorSets.txt", false);
        }


    }
    // Checks for configs, adds example and empty config
	public void checkForBonus(Player ply) {
        final Player player = ply;
        getServer().getScheduler().scheduleSyncDelayedTask(this, new Runnable(){
        
            @Override
            public void run() {
                ArmorSet added = new ArmorSet();
                for (ArmorSet set : armorSets) {
                    if (playerHas(player, set)) {
                        removeBonus(player);
                        addBonus(player, set);
                        added = set;
                        break;
                    } 
                }
                if (activeBonus.containsKey(player)) {
                    if (activeBonus.get(player) != added) {
                        removeBonus(player);
                        
                    }
                }
            }
        }, 1);
	}

    private boolean playerHas(Player player, ArmorSet Aset) {
        boolean yeh = true;
        Material[] set = Aset.armor;
        ItemStack[] playerArmor = player.getInventory().getArmorContents();
        for (int i = 0; i < 4; i++) {
            if (set[3-i] != Material.AIR) {
                if (!(playerArmor[i] != null && playerArmor[i].getType() == set[3-i])) {
                    yeh = false;
                }
            }
        }
        
        return yeh;
    }

    public void removeBonus(Player player) {
        if (activeBonus.containsKey(player)) {
            ArmorSet activeSet = activeBonus.get(player);
            PotionEffect[] effects = activeSet.effects;
            player.sendMessage("You lost armor set bonus: "+activeSet.name.split("_")[0]);
            for (PotionEffect pE : effects) {
                player.removePotionEffect(pE.getType());
            }
            activeBonus.remove(player);
            }
    }

    private void addBonus(Player player, ArmorSet set) {
        PotionEffect[] effects = set.effects;
        player.sendMessage("You got an armor set bonus: "+set.name.split("_")[0]);
        for (PotionEffect pE : effects) {
            player.addPotionEffect(pE);
        }
        activeBonus.put(player, set);
    }


}
