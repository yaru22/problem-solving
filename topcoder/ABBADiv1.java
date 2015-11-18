public class ABBADiv1 {
  private boolean _canObtain(String initial, String target) {
    if (initial.length() >= target.length()) {
      return initial.equals(target);
    }
    String planA = initial + "A";
    String planB = new StringBuffer(initial + "B").reverse().toString();
    boolean can = false;
    if (target.indexOf(planA) != -1 || new StringBuffer(target).reverse().toString().indexOf(planA) != -1) {
      can = can || _canObtain(planA, target);
    }
    if (target.indexOf(planB) != -1 || new StringBuffer(target).reverse().toString().indexOf(planB) != -1) {
      can = can || _canObtain(planB, target);
    }
    return can;
  }

  public String canObtain(String initial, String target) {
    return _canObtain(initial, target) ? "Possible" : "Impossible";
  }

  public static void main(String[] args) {
    ABBADiv1 abba = new ABBADiv1();
    System.out.println(abba.canObtain("A", "ABBA"));
  }
}
