package DesignPattern.Combined.MVC.Eg2;

public class HeartTest {
    public static void main (String[] args) {
        HeartModel heartModel = new HeartModel();
        ControllerInterface model = new HeartController(heartModel);
    }
}
