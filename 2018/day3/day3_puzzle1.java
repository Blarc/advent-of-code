import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class day3_puzzle1 {
	public static void main(String[] args) throws IOException {
		
		byte[][] fabric = new byte[1000][1000];
			
		
		if (args.length < 1) {
			System.out.println("manjka vhod");
			System.exit(1);
		}
		
		BufferedReader br = new BufferedReader(new FileReader(args[0]));
		
		Pattern p = Pattern.compile("-?\\d+");
		Matcher m;

		int result = 0;
		
		String readLine;
		for (int i = 0; (readLine = br.readLine()) != null; i++) {
			List<String> allMatches = new ArrayList<String>();
			m = p.matcher(readLine);
			while (m.find()) {
				allMatches.add(m.group());
			}
			//System.out.println(allMatches);
			
			
			int x = Integer.parseInt(allMatches.get(1));
			int y = Integer.parseInt(allMatches.get(2));
			int w = Integer.parseInt(allMatches.get(3));
			int h = Integer.parseInt(allMatches.get(4));
			
			for (int a = x; a < x+w; a++) {
				for (int b = y; b < y+h; b++) {
					if (fabric[a][b] == 1) {
						result++;
					}
					
					fabric[a][b]++;
				}
			}
			
		}
		
		System.out.println("Result: " + result);
		
	}
}
