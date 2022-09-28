import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.util.*;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class day3_puzzle2 {
	public static void main(String[] args) throws IOException {
		
		int[][] fabric = new int[1000][1000];
		boolean[] id = new boolean[1500];
		
		if (args.length < 1) {
			System.out.println("manjka vhod");
			System.exit(1);
		}
		
		BufferedReader br = new BufferedReader(new FileReader(args[0]));
		
		Pattern p = Pattern.compile("-?\\d+");
		Matcher m;

		int result = 0;
		
		String readLine;
		for (int i = 1; (readLine = br.readLine()) != null; i++) {
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
					int atmId = fabric[a][b];
					if (atmId > 0) {
						id[atmId] = true;
						id[i] = true;
					}
					fabric[a][b] = i;
				}
			}
			
		}
		
		/*for (int i = 0; i < fabric.length; i++) {
			System.out.println(Arrays.toString(fabric[i]));
		}*/
		
		
		for (int i = 1; i < id.length; i++) {
			if (!id[i]) {
				System.out.println(i);
				break;
			}
		}
		
		//System.out.println("Result: " + result);
		
	}
}
