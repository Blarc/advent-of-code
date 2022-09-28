import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class day6_puzzle2 {
	
	public static class Player  {
		int id = -1;
		int[] coord = new int[2];
		int points = 0;
		
		public Player(int id, int x, int y) {
			this.id = id;
			this.coord[0] = x;
			this.coord[1] = y;
		}
	}
	
	public static class Point {
		int owner = -1;
		int dist = -1;
		int val = 1;
		
		public Point(int owner, int dist, int val) {
			this.owner = owner;
			this.dist = dist;
			this.val = val;
		}
	}
	
	public static int abs(int a) {
		return (a > 0) ? a : -a;
	}
	
	public static int mhat(int[]a, int x, int y) {
		return abs(a[0] - x) + abs(a[1] - y);
	}
	
	public static void main(String[] args) throws IOException {
		
		int maxX = Integer.MIN_VALUE, maxY = Integer.MIN_VALUE;
		int distance = 10000;
		
		if (args.length < 1) {
			System.out.println("manjka vhod");
			System.exit(1);
		}
		
		BufferedReader br = new BufferedReader(new FileReader(args[0]));
		
		Player[] players = new Player[50];
		
		String readLine;
		for (int i = 0; (readLine = br.readLine()) != null; i++) {
			String[] coord = readLine.split(", ");
			int x = Integer.parseInt(coord[0]);
			int y = Integer.parseInt(coord[1]);
			
			maxX = (x > maxX) ? x : maxX;
			maxY = (y > maxY) ? y : maxY;
			
			players[i] = new Player(i, x, y);

		}
		
		int res = 0;
		for (int i = 0; i < maxX; i++) {
			for (int j = 0; j < maxY; j++) {
				int sum = 0;
				for (int k = 0; k < players.length; k++) {
					sum += mhat(players[k].coord, i, j);
				}
				if (sum < distance) {
					res++;
				}
			}
		}
		
		System.out.println(res);
		br.close();
	}
}
