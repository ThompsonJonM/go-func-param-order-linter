package p

func thisShouldFail(stub string, key string) {} // want "parameters:"

func thisShouldPass(key string, user string) {}

func thisHasInts(kmInt string, userInt string) {}

func thisHasMix(sub string, kmInt string, userInt string, decimal int) {} // want "parameters:"

func thisHasLog(sub string, log string) {}

func thisHasClient(Client string, sub string, user string) {}
