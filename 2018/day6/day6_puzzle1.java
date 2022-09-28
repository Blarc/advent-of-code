import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;

public class day6_puzzle1 {
	
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
		
		Point[][] map = new Point[maxX+1][maxY+1];
		for (int i = 0; i < map.length; i++) {
			for (int j = 0; j < map[i].length; j++) {
				if (i == 0 || j == 0 || i == map.length-1 || j == map[i].length-1) {
					map[i][j] = new Point(-1, -1, 10000);
				}
				else {
					map[i][j] = new Point(-1, -1, 1);
				}
			}
		}
		
		for (int k = 0; k < players.length; k++) {
			Player player = players[k];
			for (int i = 0; i < map.length; i++) {
				for (int j = 0; j < map[i].length; j++) {
					Point atm = map[i][j];
					if (atm.owner == -1) {
						atm.owner = player.id;
						atm.dist = mhat(player.coord, i, j);
						player.points += atm.val;
					}
					else if (atm.dist > mhat(player.coord, i, j)) {
						if (atm.owner != -2) {
							players[atm.owner].points -= atm.val;
						}
						atm.owner = player.id;
						atm.dist = mhat(player.coord, i, j);
						player.points += atm.val;
					}
					else if (atm.dist == mhat(player.coord, i, j)) {
						if (atm.owner != -2) {
							players[atm.owner].points -= atm.val;
						}
						atm.owner = -2;
					}
				}
			}
		}
		
		for (int i = 0; i < players.length; i++) {
			System.out.println(players[i].points);
		}
		
		br.close();
	}
}
