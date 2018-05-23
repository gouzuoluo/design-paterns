package DesignPattern.Combining.Duck.Eg3Decorator;

public class MallardDuck implements Quackable {
 
	public void quack() {
		System.out.println("Quack");
	}
 
	public String toString() {
		return "Mallard Duck";
	}
}
