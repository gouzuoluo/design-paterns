package DesignPattern.Combining.Duck.Eg2Adapter;

public class DecoyDuck implements Quackable {
	public void quack() {
		System.out.println("<< Silence >>");
	}
}
