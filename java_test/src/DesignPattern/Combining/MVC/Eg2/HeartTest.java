package DesignPattern.Combining.MVC.Eg2;

public class HeartTest {
    public static void main (String[] args) {
        HeartModel heartModel = new HeartModel();
        ControllerInterface model = new HeartController(heartModel);
    }
}
