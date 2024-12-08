async function deleteFunction(record) {
    let text = "Do you confirm the deletion of this record?";

    if (confirm(text) == true) {
        _id =
            record.parentNode.parentNode.firstChild.nextElementSibling.getElementsByTagName(
                "p"
            )[0].innerHTML;
        console.log(_id);
        const response = await fetch(`/deleteMaintainer`, {
            method: "DELETE",
            headers: {
                "Content-Type": "application/json",
                Accept: "application/json",
            },
            body: JSON.stringify({ _id }),
        });
        const message = await response.json();
        console.log(message);
    }
}
