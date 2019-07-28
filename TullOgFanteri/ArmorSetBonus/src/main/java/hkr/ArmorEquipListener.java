package hkr;

import org.bukkit.event.EventHandler;
import org.bukkit.event.Listener;
import org.bukkit.event.player.PlayerJoinEvent;
import org.bukkit.event.player.PlayerQuitEvent;

import hkr.ArmorEquipEventFiles.ArmorEquipEvent;

class ArmorEquipListener implements Listener {
    ArmorSetBonusMain pl;
    public ArmorEquipListener(ArmorSetBonusMain pl){
        this.pl = pl;
    }

    @EventHandler
    public void equipListen(ArmorEquipEvent event) {
        if (event.getPlayer().hasPermission("armorsetbonus.receive")) {
            pl.checkForBonus(event.getPlayer());
        }
    }
    @EventHandler
    public void onLogin(PlayerJoinEvent event){
        if (event.getPlayer().hasPermission("armorsetbonus.receive")) {
            pl.checkForBonus(event.getPlayer());
        }
    }
    
    // Removes bonuses when a player leaves
    // Since effects have infinite duration it must be removed incase the plugin is suddenly broken / removed
    @EventHandler
    public void onQuit(PlayerQuitEvent event) {
        pl.removeBonus(event.getPlayer());
    }


}