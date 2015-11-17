import java.util.HashSet;

public class UserName {
  public String newMember(String[] existingNames, String newName) {
    HashSet<String> names = new HashSet<String>();
    for (String name : existingNames) {
      names.add(name);
    }
    if (!names.contains(newName)) {
      return newName;
    }
    int i = 1;
    String nextName;
    do {
      nextName = newName + i;
      i++;
    } while (names.contains(nextName));
    return nextName;
  }

  public static void main(String[] args) {
    UserName userName = new UserName();
    String[] existingNames = {"MasterOfDisaster", "DingBat", "TygerTyger", "Orpheus", "WolfMan", "MrKnowItAll"};
    System.out.println(userName.newMember(existingNames, "TygerTyger"));
  }
}
