function addModal() {
    const newUser = `<div class="new-bg">
        <div class="new-modal">
            <form>
                <div class="input-group">
                    <div class = "inp-col">
                        <div class="input-field">
                            <p>Driver ID</p>
                            <input id="driver_id" type="text" required/>
                            <p id="result">Driver: John Doe<\p>
                        </div>
                        <div class="input-field">
                            <p>Status</p>
                            <select name="status" id="inp" required>
                            <option>active</option>
                            <option>cancelled</option>
                            <option>completed</option>
                            </select>
                        </div>
                        <div class="input-field">
                            <p>Time in minutes</p>
                            <input id="time" type="number" required/>
                        </div>
                    </div>

                    <div class = "inp-col">
                        <div class="input-field">
                            <p>Start Location</p>
                            <input id="start_loc"  type="text" required />
                        </div>
                        <div class="input-field">
                            <p>End Location</p>
                            <input id="end_loc" type="text" required />
                        </div>
                        <div class="input-field">
                            <p>Distance in km.</p>
                            <input id="distance" type="number" required />
                        </div>
                    </div>

                    
                </div>
                <div class="textarea-field">
                        <p>Description</p>
                        <textarea  id="description" type="text"></textarea>
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

async function getDrivers() {
    let response_driver = await fetch(`/drivers`, {
        headers: {
            "Content-type": "application/json; charset=UTF-8",
            Accept: "application/json",
        },
    });
    return await response_driver.json();
}

const $addButton = document.querySelector(".add-button");
const $app = document.querySelector("#app");

let driverName = document.getElementById("result");

$addButton.onclick = async function () {
    addModal();
    const checkDrivers = getDrivers();
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

    const end_loc = document.querySelector("#end_loc");

    const time = document.querySelector("#time");
    const status = document.getElementById("inp");
    const start_loc = document.querySelector("#start_loc");
    const distance = document.querySelector("#distance");
    const description = document.querySelector("#description");

    function isEmpty(value) {
        return value.trim() === "";
    }

    save.onclick = async function () {
        if (
            isEmpty(end_loc.value) ||
            isEmpty(driver_id.value) ||
            isEmpty(time.value) ||
            isEmpty(status.value) ||
            isEmpty(start_loc.value) ||
            isEmpty(distance.value)
        ) {
            text = "ivalid input";
            alert(text);
        } else {
            const response1 = await fetch("/createRoute", {
                method: "POST",
                body: JSON.stringify({
                    time_for_route: parseInt(time.value),
                    start_location: start_loc.value,
                    end_location: end_loc.value,
                    description: description.value,
                    driver_id: driver_id.value,
                    distance: parseInt(distance.value),
                    status: status.value,
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
