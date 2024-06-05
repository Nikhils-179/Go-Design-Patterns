package Files.Java;
public class Mango {

    private static Mango mango = new Mango();

    public static Mango getMango(){
        return mango;
    }
}

//Eager way of creating singleton object , instance already created

