package hkr.playtimeTracker;


public class Tuple implements Comparable{ 
  public final String x; 
  public final long y; 
  public Tuple(String x, long y) { 
    this.x = x; 
    this.y = y; 
  }

    @Override
    public int compareTo(Object obj) {
        if (obj instanceof Tuple) {
           Tuple o = (Tuple) obj; 
           return Long.compare(y, o.y);
        } else {
            return 0;
        }
            
    }
  
} 