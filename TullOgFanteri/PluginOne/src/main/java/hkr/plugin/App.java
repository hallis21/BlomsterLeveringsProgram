package hkr.plugin;

import java.io.File;

import org.bukkit.plugin.java.JavaPlugin;


/**
 * Hello world!
 *
 */
public class App extends JavaPlugin
{
    public final PlaytimeListener playtimeL = new PlaytimeListener(this);
    @Override
    public void onEnable() {
        File f = new File(this.getDataFolder() + "/");
        if (!f.exists()){
            f.mkdir();
        }

        
        getServer().getPluginManager().registerEvents(playtimeL, this);
        getServer().getScheduler().scheduleSyncRepeatingTask(this, new Runnable(){
            
            @Override
            public void run() {
                playtimeL.updateFile();
            }
        }, 1000, 6000);
        
        new CommandInit(this).InitCommands();
        
        getLogger().info("Loaded.");
    }

    @Override
    public void onDisable() {
        playtimeL.updateFile();
        getLogger().info("Shutting down.");
    }
}