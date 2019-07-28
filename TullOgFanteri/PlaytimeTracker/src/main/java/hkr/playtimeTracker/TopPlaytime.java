package hkr.playtimeTracker;

import java.util.ArrayList;
import java.util.Objects;
import java.util.concurrent.TimeUnit;

import org.bukkit.command.Command;
import org.bukkit.command.CommandExecutor;
import org.bukkit.command.CommandSender;
import org.bukkit.entity.Player;

public class TopPlaytime implements CommandExecutor {
    App pl;
    public TopPlaytime(App pl) {
        this.pl = pl;
    }

    @Override
    public boolean onCommand(CommandSender sender, Command command, String label, String[] args) {
        Player player = (Player) sender;
        ArrayList<Tuple> players = pl.playtimeL.getTop();
        // Prints list of top players
        int i = 1;
        player.sendMessage("Top players (playtime):");
        for (Tuple tup : players) {
            if (i>10) {
                break;
            }
            long timePlayed = tup.y;
            String days = Objects.toString(TimeUnit.MILLISECONDS.toDays(timePlayed)); 
            String hours = Objects.toString(TimeUnit.MILLISECONDS.toHours(timePlayed) - TimeUnit.MILLISECONDS.toHours(TimeUnit.MILLISECONDS.toDays(timePlayed)));
            String mins = Objects.toString(TimeUnit.MILLISECONDS.toMinutes(timePlayed) -  TimeUnit.HOURS.toMinutes(TimeUnit.MILLISECONDS.toHours(timePlayed)));
            String secs = Objects.toString(TimeUnit.MILLISECONDS.toSeconds(timePlayed) -  TimeUnit.MINUTES.toSeconds(TimeUnit.MILLISECONDS.toMinutes(timePlayed)));
  


            player.sendMessage(i+". "+tup.x+": "+ days+" days, "+hours+" hours, "+mins+" minutes, " + secs + " seconds.");
            i++;
        }
        return true;
	}
    
}