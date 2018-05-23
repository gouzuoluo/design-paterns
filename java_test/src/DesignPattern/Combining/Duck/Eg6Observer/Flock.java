package DesignPattern.Combining.Duck.Eg6Observer;

import java.util.ArrayList;
import java.util.Iterator;

public class Flock implements Quackable {
	ArrayList<Quackable> ducks = new ArrayList<Quackable>();
 
	public void add(Quackable quacker) {
		ducks.add(quacker);
	}
 
	public void quack() {
		Iterator<Quackable> iterator = ducks.iterator();
		while (iterator.hasNext()) {
			Quackable quacker = iterator.next();
			quacker.quack();
		}
	}

	public void registerObserver(Observer observer) {
		Iterator<Quackable> iterator = ducks.iterator();
		while (iterator.hasNext()) {
			Quackable duck = (Quackable)iterator.next();
			duck.registerObserver(observer);
		}
	}

	public void notifyObservers() { }
	public String toString() {
		return "Flock of Quackers";
	}
}
