package days;

func GetDayFunctions() map[int]func() {
    dayFuncs := map[int]func(){
        1: Day1,
    };

    return dayFuncs;
}
