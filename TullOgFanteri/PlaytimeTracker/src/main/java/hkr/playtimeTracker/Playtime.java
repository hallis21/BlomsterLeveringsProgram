package hkr.playtimeTracker;

import java.util.Objects;
import java.util.concurrent.TimeUnit;

import org.bukkit.command.Command;
import org.bukkit.command.CommandExecutor;
import org.bukkit.command.CommandSender;
import org.bukkit.entity.Player;

public class Playtime implements CommandExecutor {

    App pl;
    PlaytimeListener pL;
    public Playtime(App pl){
      this.pl = pl;
      this.pL = pl.playtimeL;
    }
    @Override
    public boolean onCommand(CommandSender sender, Command command, String label, String[] args) {
      Player player = (Player) sender;

      try {
        long timePlayed = pL.getPlayed(player);
        
        String days = Objects.toString(TimeUnit.MILLISECONDS.toDays(timePlayed)); 
        String hours = Objects.toString(TimeUnit.MILLISECONDS.toHours(timePlayed) - TimeUnit.MILLISECONDS.toHours(TimeUnit.MILLISECONDS.toDays(timePlayed)));
        String mins = Objects.toString(TimeUnit.MILLISECONDS.toMinutes(timePlayed) -  TimeUnit.HOURS.toMinutes(TimeUnit.MILLISECONDS.toHours(timePlayed)));
        String secs = Objects.toString(TimeUnit.MILLISECONDS.toSeconds(timePlayed) -  TimeUnit.MINUTES.toSeconds(TimeUnit.MILLISECONDS.toMinutes(timePlayed)));
  
        player.sendMessage("Time played: "+ days+" days, "+hours+" hours, "+mins+" minutes, " + secs + " seconds.");
  
        return true;
        
      } catch (Exception e) {
        return onCommand(sender, command, label, args);
      }
	}
    
}