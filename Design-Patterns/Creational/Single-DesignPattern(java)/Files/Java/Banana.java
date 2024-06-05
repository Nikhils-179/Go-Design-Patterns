package Files.Java;
import java.io.Serializable;

public class Banana implements Serializable{
    
    public static Banana banana = new Banana();
    private Banana(){

    }

    public static Banana getBanana(){

        synchronized(Banana.class){
            if (banana == null) {
            return banana;
        }
        return banana;
    }
    }
}
//If code requires thread safety
//Two threads running t1,t2.

