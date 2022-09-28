import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;


public class day4_puzzle2 {
	
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
		int[][] guardRes = new int[4000][60];

		Event temp = first.next;
		while (temp != null) {
			System.out.println(temp.msg);
			
			String[] msgSplit = temp.msg.split(" ");
			if (msgSplit[2].equals("Guard")) {
				guardId = Integer.parseInt(msgSplit[3].split("#")[1]);
			}
			
			else if (msgSplit[2].equals("falls")) {
				startSleep[0] = temp.hour;
				startSleep[1] = temp.minute;
			}
			
			else if (msgSplit[2].equals("wakes")) {
				int timeSlept = (temp.hour * 60 + temp.minute) - (startSleep[0] * 60 + startSleep[1]);
				for (int i = startSleep[1]; i < startSleep[1] + timeSlept; i++) {
					guardRes[guardId][i]++;
				}
				
			}
			//System.out.println(temp.msg.split(" ")[2]);
			temp = temp.next;
		}
		
		
		
		for (int i = 0; i < guardRes.length; i++) {
			int max = -1;
			int bestId = -1;
			for (int j = 0; j < guardRes[i].length; j++) {
				if (guardRes[i][j] > max) {
					max = guardRes[i][j];
					bestId = j;
				}
			}
			
			if (bestId > 0) {
				//System.out.println(Arrays.toString(guardRes[i]));
				System.out.println(i + " " + bestId + " " + max);
			}
			
		}
	}
}
