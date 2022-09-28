import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class day4_puzzle1 {
	
	public static class Event {
		int month;
		int day;
		int hour;
		int minute;
		String msg;
		
		Event next;
		
		public Event (int month, int day, int hour, int minute, String msg) {
			this.month = month;
			this.day = day;
			this.hour = hour;
			this.minute = minute;
			this.msg = msg;
		}
		
		public Event (int month, int day, int hour, int minute, String msg, Event next) {
			this.month = month;
			this.day = day;
			this.hour = hour;
			this.minute = minute;
			this.msg = msg;
			this.next = next;
		}
		
		public boolean isLessThan(Event a) {
			if (month < a.month) {
				return true;
			}
			else if (month > a.month) {
				return false;
			}
			else {
				if (day < a.day) {
					return true;
				}
				else if (day > a.day) {
					return false;
				}
				else {
					if (hour < a.hour) {
						return true;
					}
					else if (hour > a.hour) {
						return false;
					}
					else {
						if (minute < a.minute) {
							return true;
						}
						else if (minute > a.minute) {
							return false;
						}
					}
				}
			}
			return false;
		}
	}
	

	public static void main(String[] args) throws IOException {

		if (args.length < 1) {
			System.out.println("manjka vhod");
			System.exit(1);
		}
		
		BufferedReader br = new BufferedReader(new FileReader(args[0]));
		
		Event[] events = new Event[1128];
		
		Pattern p = Pattern.compile("\\d+");
		Matcher m;
		
		Event first = new Event(0, 0, 0, 0, "0 0");
		String readLine;
		for (int i = 0; (readLine = br.readLine()) != null; i++) {
			//System.out.println(i);
			List<String> allMatches = new ArrayList<String>();
			m = p.matcher(readLine);
			while (m.find()) {
				allMatches.add(m.group());
			}
			
			int month = Integer.parseInt(allMatches.get(1));
			int day = Integer.parseInt(allMatches.get(2));
			int hour = Integer.parseInt(allMatches.get(3));
			int minute = Integer.parseInt(allMatches.get(4));
			
			Event newEvent = new Event(month, day, hour, minute, readLine);
			
			Event temp = first;
			
			while (temp.next != null) {
				if (temp.next.isLessThan(newEvent)) {
					temp = temp.next;
				}
				else {
					break;
				}
			}
			
			newEvent.next = temp.next;
			temp.next = newEvent;
			
			
		}
		
		int guardId = 0;
		int[] startSleep = new int[2];
		int[][] guardRes = new int[4000][2];

		
		Event temp = first.next;
		while (temp != null) {
			//System.out.println(temp.msg);
			
			String[] msgSplit = temp.msg.split(" ");
			if (msgSplit[2].equals("Guard")) {
				guardId = Integer.parseInt(msgSplit[3].split("#")[1]);
			}
			
			else if (msgSplit[2].equals("falls")) {
				startSleep[0] = temp.hour;
				startSleep[1] = temp.minute;
			}
			
			else if (msgSplit[2].equals("wakes")) {
				int timeSlept = (temp.hour * 60 + temp.minute) - (startSleep[0] * 60 + startSleep[1])-1;
				//System.out.println(timeSlept);
				guardRes[guardId][0] += timeSlept;
				if (guardRes[guardId][1] == 0 || guardRes[guardId][1] < timeSlept) {
					guardRes[guardId][1] = timeSlept;
				}
			}
			//System.out.println(temp.msg.split(" ")[2]);
			temp = temp.next;
		}
		
		int max = -1;
		int bestId = 0;
		
		for (int i = 0; i < guardRes.length; i++) {
			if (guardRes[i][0] > max) {
				max = guardRes[i][0];
				bestId = i;
			}
		}
		
		System.out.println(bestId);
		
		int[] sleepCount = new int[60];
		int startSleeping = 0;
		
		temp = first.next;
		while (temp != null) {
			
			String[] msgSplit = temp.msg.split(" ");
			if (msgSplit[2].equals("Guard")) {
				guardId = Integer.parseInt(msgSplit[3].split("#")[1]);
			}
			
			if (guardId == bestId) {
				//System.out.println(temp.msg);
				
				if (msgSplit[2].equals("falls")) {
					startSleep[0] = temp.hour;
					startSleep[1] = temp.minute;
				}
				
				else if (msgSplit[2].equals("wakes")) {
					//System.out.println(startSleep[1]);
					int timeSlept = (temp.hour * 60 + temp.minute) - (startSleep[0] * 60 + startSleep[1]);
					//System.out.println(timeSlept);
					for (int i = startSleep[1]; i < startSleep[1] + timeSlept; i++) {
						sleepCount[i]++;
					}
				}
			}
			temp = temp.next;
		}
		
		max = -1;
		int bestIndex = -1;
		for (int i = 0; i < sleepCount.length; i++) {
			if (sleepCount[i] > max) {
				max = sleepCount[i];
				bestIndex = i;
			}
		}
		
		System.out.println(bestIndex);
		System.out.println(bestId * bestIndex);
	}
}
