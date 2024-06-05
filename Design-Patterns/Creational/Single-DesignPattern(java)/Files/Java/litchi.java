package Files.Java;
public class litchi {
    private static litchi litchi;
    //constructor
    private litchi(){
    }
    public static litchi getlitchi(){
        if(litchi == null) {
            litchi=  new litchi();
        }
        return litchi;
    }
    
}

//lazy way of creating object : when needed create else donot
