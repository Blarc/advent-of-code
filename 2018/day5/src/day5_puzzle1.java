import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;


public class day5_puzzle1 {
	
	public static class StackElement
	{
		Object element;
		StackElement next;

		StackElement()
		{
			element = null;
			next = null;
		}
	}

	public static class Stack
	{
		//StackElement -> StackElement -> StackElement -> ... -> StackElement
		//     ^
		//     |
		//    top                                                   
		//
		// elemente dodajamo in brisemo vedno na zacetku seznama (kazalec top)
		
		
		private StackElement top;
		
		public Stack()
		{
			makenull();
		}
		
		public void makenull()
		{
			top = null;
		}
		
		public boolean empty()
		{
			return (top == null);
		}
		
		public Object top()
		{
			if (!empty())
				return top.element;
			else
				return null;
		}
		
		public void push(Object obj)
		{
			StackElement el = new StackElement();
			el.element = obj;
			el.next = top;
			
			top = el;
		}
		
		public void pop()
		{
			if (!empty())
			{
				top = top.next;
			}
		}
	}

	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new FileReader(args[0]));
		
		int c, stackedChar, countRemoved = 0, countAll = -1;
		
		
		Stack stack = new Stack();
		
		
		
		while ((c = br.read()) != -1) {
			countAll++;
			if (!stack.empty()) {
				stackedChar = (int)stack.top();
				
				System.out.println((char)stackedChar + " " + (char)c);
				System.out.println(stackedChar - c);
				if (stackedChar - c == 32 || stackedChar - c == -32) {
					countRemoved += 2;
					stack.pop();
				}
				else {
					stack.push(c);
				}
			}
			else {
				stack.push(c);
			}
		}
		
		System.out.println("All: " + countAll);
		System.out.println("Removed: " + countRemoved);
		System.out.println("Result: " + (countAll - countRemoved));
		
		
		
		br.close();
		
	}

}
