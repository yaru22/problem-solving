public class BettingMoney {
  public int moneyMade(int[] amounts, int[] centsPerDollar, int finalResult) {
    int moneyMade = 0;
    for (int i = 0; i < amounts.length; i++) {
      if (i == finalResult) {
        moneyMade -= (amounts[i] * centsPerDollar[i]);
      } else {
        moneyMade += (amounts[i] * 100);
      }
    }
    return moneyMade;
  }

  // public static void main(String[] args) {
  //   BettingMoney betting = new BettingMoney();
  //   int[] amounts = { 10, 20, 30 };
  //   int[] centsPerDollar = { 20, 30, 40 };
  //   System.out.println(betting.moneyMade(amounts, centsPerDollar, 1));
  // }
}
