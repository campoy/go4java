public class Programmer {
    private int caffeineLevel;

    public void drinkCoffee() {
        this.caffeineLevel++;
    }

    public bool isAwake() {
        return this.caffeineLevel > 42;
    }
}