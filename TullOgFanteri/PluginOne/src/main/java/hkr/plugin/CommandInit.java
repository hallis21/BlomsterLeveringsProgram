package hkr.plugin;


public class CommandInit {
    App pl;
    public CommandInit(App pl){
        this.pl = pl;
    }

	public void InitCommands() {
        played();
        topPlayed();
	}
    
    private void played() {
        pl.getCommand("played").setExecutor(new Playtime(pl));
    }
    
    private void topPlayed() {
        pl.getCommand("playedtop").setExecutor(new TopPlaytime(pl));
    }

}