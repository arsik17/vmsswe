async function getAllMemberss() {
    const response = await fetch("/reports", {
        headers: {
            "content-type": "application/json",
            accept: "application/json",
        },
    });

    responseBody = await response.json();

    console.log(responseBody);

    const $list = document.querySelector(".list");
    let elements = "";
    responseBody.forEach((oneinstance) => {
        elements += `
        <div class="element">
            <div style="width: 135px;" class="text"  id="member_user_id">
                <p>${oneinstance._id}</p>
            </div>
            <div style="width: 135px;" class="text"  id="route_id">
                <p>${oneinstance.route_id}</p>
            </div>
            <div style="width: 135px;" class="text"  id="vehicle_id">
                <p>${oneinstance.route_id}</p>
            </div>
            <div style="flex: 1;" class="text">
                <p>${oneinstance.vehicle_id}</p>
            </div>
            <div style="flex: 1.2;" class="text">
                <p>${oneinstance.total_distance} km.</p>
            </div>
            <div style="flex: 0.9;" class="text">
                <p>${oneinstance.money_spent} $</p>
            </div>
            <div style="flex: 0.9;" class="text">
                <p>${oneinstance.fuel_usage} L</p>
            </div>
            <div class = "change-bt">

            <button onclick="deleteFunction(this)" type="button" class="change" id="d">
                <img
                    src="/assets/bin.png"
                    alt=""
                />
                </button>
            <button onclick="viewFunction(this)" type="button" class="change" id="c">
                <img
                    src="/assets/refresh.png"
                    alt=""
                />
                    </button>
            </div>

        </div>`;
        $list.innerHTML = elements;
    });
}
getAllMemberss();

function element(id) {
    return `<div class="element" id="${id}">
    </div>
    `;
}

function id(content) {
    return `
        <div class="text">
            <p>${content}</p>
        </div>
    `;
}
function email(content) {
    return `
        <div class="text">
            <p>${content}</p>
        </div>
    `;
}
function given_name(content) {
    return `
        <div class="text">
            <p>${content}</p>
        </div>
    `;
}
function surname(content) {
    return `
        <div class="text">
            <p>${content}</p>
        </div>
    `;
}

function city(content) {
    return `
        <div class="text">
            <p>${content}</p>
        </div>
    `;
}

function phone_number(content) {
    return `
        <div class="text">
            <p>${content}</p>
        </div>
    `;
}

function change() {
    return `
        <button type="button" id="change">
        <img
            src="/assets/refresh.png"
            alt=""
        />
            </button>
    `;
}
