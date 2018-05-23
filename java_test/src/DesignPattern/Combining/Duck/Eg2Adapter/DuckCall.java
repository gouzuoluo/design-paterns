package DesignPattern.Combining.Duck.Eg2Adapter;

public class DuckCall implements Quackable {
	public void quack() {
		System.out.println("Kwak");
	}
}
