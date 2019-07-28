package hkr.playtimeTracker;

import java.io.File;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.FileWriter;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.StandardCopyOption;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Iterator;
import java.util.List;
import java.util.Set;

import org.bukkit.entity.Player;
import org.bukkit.event.EventHandler;
import org.bukkit.event.Listener;
import org.bukkit.event.player.PlayerJoinEvent;
import org.bukkit.event.player.PlayerQuitEvent;
import org.json.simple.JSONArray;
import org.json.simple.JSONObject;
import org.json.simple.parser.JSONParser;
import org.json.simple.parser.ParseException;

public class PlaytimeListener implements Listener {
    private ArrayList<Player> playerOnline = new ArrayList<>();
    private App pl;
    private File fil;

    public PlaytimeListener(App pl){
        this.pl = pl;
        File f = new File(pl.getDataFolder() + "/");
        if (!f.exists()) {
            f.mkdir();
        }
        fil = new File(pl.getDataFolder() + "/player.json");
        if (!fil.exists()) {
            try {
                fil.createNewFile();
                
            } catch (IOException e) {
                e.printStackTrace();
            }

        }
        try {
            JSONObject jObj = (JSONObject) new JSONParser().parse(new FileReader(fil));
            pl.getLogger().info("Loaded logs");

        } catch (IOException | ParseException e) {
            fixCorruptedFile();

		}
        
    }

    // Checks when a player joins and starts logging time
    @EventHandler
    public void onPlayerJoin(PlayerJoinEvent event) {
        Player player = event.getPlayer();
        playerUpdate(player);

    }

    private void playerUpdate(Player player) {
        try {
            JSONObject jObj;
            jObj = (JSONObject) new JSONParser().parse(new FileReader(fil));
            if (isLogged(player)) {
                JSONArray jArr = new JSONArray();
                JSONArray tempArray = (JSONArray) jObj.get(player.getName());

                jArr.add(tempArray.get(0));
                jArr.add(System.currentTimeMillis());
                jObj.replace(player.getName(), jArr);

                FileWriter fR = new FileWriter(fil);
                fR.write(jObj.toJSONString());
                fR.flush();
            } else {
                JSONArray jArr = new JSONArray();
                jArr.add(0);
                jArr.add(System.currentTimeMillis());
                jObj.put(player.getName(), jArr);
                FileWriter fR = new FileWriter(fil);
                fR.write(jObj.toJSONString());
                fR.flush();
            }
        } catch (IOException | ParseException e) {
            fixCorruptedFile();
            playerUpdate(player);
        }
        playerOnline.add(player);
    }
    @EventHandler
    public void onPlayerLeave(PlayerQuitEvent event) {
        Player player = event.getPlayer();
        try {
            JSONObject jObj;
            jObj = (JSONObject) new JSONParser().parse(new FileReader(fil));
            JSONArray jArr = new JSONArray();
            JSONArray tempArray = (JSONArray) jObj.get(player.getName());
            long timePlayed = (System.currentTimeMillis() - (long) tempArray.get(1))+(long) tempArray.get(0);
            jArr.add(timePlayed);
            jArr.add(System.currentTimeMillis());
            jObj.replace(player.getName(), jArr);

            FileWriter fR = new FileWriter(fil);
            fR.write(jObj.toJSONString());
            fR.flush();       
        } catch (IOException | ParseException e) {
            fixCorruptedFile();
            onPlayerLeave(event);
        }
        playerOnline.remove(event.getPlayer());
    }
    
    public void updateFile() {
        try {
            JSONObject jObj;
            jObj = (JSONObject) new JSONParser().parse(new FileReader(fil));
            for (Player player : playerOnline) {
                JSONArray jArr = new JSONArray();
                JSONArray tempArray = (JSONArray) jObj.get(player.getName());
                long timePlayed = (System.currentTimeMillis() - (long) tempArray.get(1)) + (long) tempArray.get(0);
                jArr.add(timePlayed);
                jArr.add(System.currentTimeMillis());
                jObj.replace(player.getName(), jArr);

            }
            FileWriter fR = new FileWriter(fil);
            fR.write(jObj.toJSONString());
            fR.flush();
            
            } catch (IOException | ParseException e) {
                fixCorruptedFile();
                updateFile();
            }
    }

    private boolean isLogged(Player ply) {
        String playerName = ply.getName();
        try {
            JSONObject jObj = (JSONObject) new JSONParser().parse(new FileReader(fil));
            if (jObj.containsKey(playerName)) {
                return true;
            } else {
                return false;
            }
        } catch (IOException | ParseException e) {
            fixCorruptedFile();
            return isLogged(ply);
        }

    }
    public long getPlayed(Player player){
        updateFile();
        try {
            JSONObject jObj;
            jObj = (JSONObject) new JSONParser().parse(new FileReader(fil));
            JSONArray tempArray = (JSONArray) jObj.get(player.getName());
            long timePlayed = (long) tempArray.get(0);
            return timePlayed;
        } catch (IOException | ParseException e) {
            fixCorruptedFile();
            return getPlayed(player);
        }
    }
    private void fixCorruptedFile(){
        pl.getLogger().info("Corrupted logs, old logs have been saved as seperate file.");
        fil = new File(pl.getDataFolder() + "/player.json");
        if (!fil.exists()) {
            try {
                fil.createNewFile();
                FileWriter fR = new FileWriter(fil);
                fR.write("{}");
                fR.close();

            } catch (IOException e) {
                e.printStackTrace();
            }

        } else {
            try {
                File nyFil = new File(pl.getDataFolder() + "/player_bad_version.json");
                if (nyFil.exists()) {
                    nyFil.delete();
                }

                Path copied = Paths.get(nyFil.getPath());
                Path originalPath = Paths.get(fil.getPath());;
                Files.copy(originalPath, copied, StandardCopyOption.REPLACE_EXISTING);

                fil.delete();
                fil.createNewFile();
                FileWriter fR = new FileWriter(fil);
                fR.write("{}");
                fR.close();
            } catch (IOException e) {
                e.printStackTrace();
                pl.getLogger().info("Big bad.");
            }
        }
        // Repopulates list with players online
        for (Player player : playerOnline) {
        playerUpdate(player);
        }
    }

	public ArrayList<Tuple> getTop() {
        updateFile();
        try {
            JSONObject jObj;
            jObj = (JSONObject) new JSONParser().parse(new FileReader(fil));
            ArrayList<Tuple> players = new ArrayList<>();
            for (Iterator iterator = jObj.keySet().iterator(); iterator.hasNext();) {
                String key = (String) iterator.next();

                JSONArray played = (JSONArray)jObj.get(key);
                players.add(new Tuple(key, (long) played.get(0)));
            }
            Collections.sort(players);
            Collections.reverse(players);
            return players;

        } catch (IOException | ParseException e) {
            fixCorruptedFile();
            return getTop();
        }

	}
}