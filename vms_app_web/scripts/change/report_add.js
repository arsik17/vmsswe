function addModal() {
    const newUser = `<div class="new-bg">
        <div class="new-modal">
            <form>
                <div class="input-group">
                    <div class = "inp-col">
                        <div class="input-field">
                            <p>Driver ID</p>
                            <input id="driver_id" type="text" required/>
                            <p id="driver_result">Driver: John Doe<\p>
                        </div>
                        <div class="input-field">
                            <p>Route ID</p>
                            <input id="route_id" type="text" required/>
                        </div>
                        <div class="input-field">
                            <p>Total Distance</p>
                            <input id="total_distance" type="number" required/>
                        </div>
                    </div>

                    <div class = "inp-col">
                        <div class="input-field">
                            <p>Vehicle ID</p>
                            <input id="vehicle_id"  type="text" required />
                            <p id="vehicle_result">LPN: 1241411<\p>
                        </div>
                        <div class="input-field">
                            <p>Fuel Usage</p>
                            <input id="fuel_usage" type="number" required />
                        </div>
                        <div class="input-field">
                            <p>Money Spent</p>
                            <input id="money_spent" type="number" required />
                        </div>
                    </div>

                    
                </div>
                    <div class="users-cancelsub">
                        <button type="submit" id="cancel" class="addnewbut">
                            <p>Cancel</p>
                        </button>
                        <button type="button"  id="save" class="addnewbut">
                            <p>Save</p>
                        </button>
                        </button>
                    </div>
            </form>
        </div>
    </div>`;
    $app.insertAdjacentHTML("afterbegin", newUser);
}

// async function getDrivers() {
//     let response_driver = await fetch(`/drivers`, {
//         headers: {
//             "Content-type": "application/json; charset=UTF-8",
//             Accept: "application/json",
//         },
//     });
//     return await response_driver.json();
// }

const $addButton = document.querySelector(".add-button");
const $app = document.querySelector("#app");
let modalAdd = "";
let driverName = document.getElementById("result");

$addButton.onclick = async function () {
    addModal();
    const driver_id = document.querySelector("#driver_id");

    // driver_id.onchange = function () {
    //     checkDrivers.forEach((element) => {
    //         if (element._id != undefined) {
    //             ready = true;
    //             driverName.innerHTML = `Driver: ${element.first_name} ${element.last_name}`;
    //         } else {
    //             ready = false;
    //             driverName.innerHTML = "No such user exists";
    //         }
    //     });
    // };

    cancel.onclick = function () {
        modalAdd = document.querySelector(".new-bg").remove();
    };

    const route_id = document.querySelector("#route_id");
    const fuel_usage = document.querySelector("#fuel_usage");
    const total_distance = document.querySelector("#total_distance");
    const vehicle_id = document.querySelector("#vehicle_id");
    const money_spent = document.querySelector("#money_spent");

    function isEmpty(value) {
        return value.trim() === "";
    }

    save.onclick = async function () {
        if (
            isEmpty(route_id.value) ||
            isEmpty(driver_id.value) ||
            isEmpty(fuel_usage.value) ||
            isEmpty(total_distance.value) ||
            isEmpty(money_spent.value)
        ) {
            text = "ivalid input";
            alert(text);
        } else {
            const response1 = await fetch("/createReport", {
                method: "POST",
                body: JSON.stringify({
                    route_id: route_id.value,
                    fuel_usage: parseFloat(fuel_usage.value),
                    total_distance: parseInt(total_distance.value),
                    vehicle_id: vehicle_id.value,
                    driver_id: driver_id.value,
                    money_spent: parseInt(money_spent.value),
                }),
                headers: {
                    "Content-type": "application/json; charset=UTF-8",
                    Accept: "application/json",
                },
            });
            const message1 = await response1.json();

            modalAdd = document.querySelector(".new-bg").remove();
        }
    };
};
