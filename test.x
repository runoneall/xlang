os := import("os")
custom := import("custom")
fmt := import("fmt")


add_result := custom.add(5, 3)
subtract_result := custom.subtract(9, 3)

fmt.println("Addition result:",add_result)
fmt.println("Subtraction result:",subtract_result)

fmt.print("Input: ")
ui := fmt.scanln()

fmt.println("You input:", ui)