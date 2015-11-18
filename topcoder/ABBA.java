public class ABBA {
  private boolean _canObtain(String initial, String target) {
    if (initial.length() == target.length()) {
      return initial.equals(target);
    }
    char last = target.charAt(target.length() - 1);
    String prev;
    if (last == 'A') {
      prev = target.substring(0, target.length() - 1);
    } else {  // 'B'
      prev = new StringBuffer(target.substring(0, target.length() - 1)).reverse().toString();
    }
    return _canObtain(initial, prev);
  }

  public String canObtain(String initial, String target) {
    return _canObtain(initial, target) ? "Possible" : "Impossible";
  }

  public static void main(String[] args) {
    ABBA abba = new ABBA();
    System.out.println(abba.canObtain("A", "ABBA"));
  }
}
