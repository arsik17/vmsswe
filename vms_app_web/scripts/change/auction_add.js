function addModal() {
    const newUser = `<div class="new-bg">
        <div class="new-modal">
            <form>
                <div class="input-group">
                    <div class = "inp-col">
                        <div class="input-field">
                            <p>Vehicle ID</p>
                            <input id="vehicle_id" type="text" required/>
                            <p id="result">VIN: 1213131<\p>
                        </div>
                        <div class="input-field">
                            <p>Status</p>
                            <select name="status" id="inp" required>
                            <option>active</option>
                            <option>sold</option>
                            <option>inactive</option>
                            </select>
                        </div>
                        
                    </div>

                    <div class = "inp-col">
                        <div class="input-field">
                            <p>Starting Price</p>
                            <input id="starting_price"  type="text" required />
                        </div>
                        <div class="input-field">
                            <p>Auction Name</p>
                            <input id="name" type="text" required/>
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
    const vehicle_id = document.querySelector("#vehicle_id");

    cancel.onclick = function () {
        modalAdd = document.querySelector(".new-bg").remove();
    };

    const name = document.querySelector("#name");
    const status = document.getElementById("inp");
    const starting_price = document.querySelector("#starting_price");
    const description = document.querySelector("#description");

    function isEmpty(value) {
        return value.trim() === "";
    }

    save.onclick = async function () {
        if (
            isEmpty(vehicle_id.value) ||
            isEmpty(name.value) ||
            isEmpty(description.value) ||
            isEmpty(starting_price.value) ||
            isEmpty(status.value)
        ) {
            text = "ivalid input";
            alert(text);
        } else {
            const response1 = await fetch("/createAuction", {
                method: "POST",
                body: JSON.stringify({
                    vehicle_id: vehicle_id.value,
                    name: name.value,
                    description: description.value,
                    starting_price: parseInt(starting_price.value),
                    current_price: parseInt(starting_price.value),
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
