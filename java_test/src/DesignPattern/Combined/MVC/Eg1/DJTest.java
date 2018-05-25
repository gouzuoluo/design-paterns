package DesignPattern.Combined.MVC.Eg1;

public class DJTest {
    public static void main(String[] args) {
        //先建立一个模型
        BeatModelInterface model = new BeatModel();

        //然后创建一个控制器，并将模型传给它。（记住，控制器创建视图，所有我们不需要把控制器介绍给视图认识）
        ControllerInterface controller = new BeatController(model);
    }
}
