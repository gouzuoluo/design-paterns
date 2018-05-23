package DesignPattern.Combining.Duck.Eg1;

public class DecoyDuck implements Quackable {
	public void quack() {
		System.out.println("<< Silence >>");
	}
}
