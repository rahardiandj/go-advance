CoffeeMachineInterface.java

public interface CoffeeMachineInterface {
    public void choseFirstSelection(){
        //implement me
    }

    public void choseSecondSelection(){
        //implement me
    }

}


OldCoffeeMachine.java

public class OldCoffeeMachine {

    public void SelectA() {
        //implement me
    }


    public void SelectB() {
        // implement me
    }


}






CoffeeTouchscreenAdapter.java

public class CoffeeTouchscreenAdapter implements CoffeeMachineInterface {

    private OldCoffeeMachine OldVendingMachine;

    public CoffeeTouchscreenAdapter(OldCoffeeMachine oldMachine) {
        this.OldVendingMachine = oldMachine
    }




}