loadAllCustomers();

//Save customer
$("#btnAdd").click(function () {

    console.log("Save button working")

    let id = $("#custID").val();
    let name = $("#custName").val();
    let address = $("#address").val();
    let nic = $("#nic").val();
    let contac = $("#contact").val();

    console.log("Contact number "+ contact)

    $.ajax({
        method: "post",
        url: "http://localhost:8000/api/customer",
        contentType: "application/json",
        data: JSON.stringify({
            "id": id,
            "name": name,
            "address": address,
            "nic": nic,
            "contact": parseInt(contac)
        }),
        success: function (res) {
            alert("Customer Registered");
            loadAllCustomers()
        },
        error: function (ob, txtStatus, error) {
            console.log(error);
            console.log(txtStatus);
            console.log(ob);
        }
    });

});


//update customer
$("#btnUpdate").click(function () {

    console.log("Update button working")

    let id = $("#custID").val();
    let name = $("#custName").val();
    let address = $("#address").val();
    let nic = $("#nic").val();
    let contac = $("#contact").val();

    console.log("working button");

    $.ajax({
        method: "put",
        url: "http://localhost:8000/api/customer/" + id,
        contentType: "application/json",
        data: JSON.stringify({
            "id": id,
            "name": name,
            "address": address,
            "nic": nic,
            "contact": parseInt(contac)
        }),
        success: function (res) {
            alert("Customer Updated");
            loadAllCustomers()
        },
        error: function (ob, txtStatus, error) {
            console.log(error);
            console.log(txtStatus);
            console.log(ob);
        }
    });

});


//delete customer
$("#btnDelete").click(function (){

    console.log("Delete button working")

    let customerID=$("#custID").val();
    $.ajax({
        method:"DELETE",
        url:"http://localhost:8000/api/customer/"+ customerID,
        success:function (res){
            alert("Customer Removed..!");
            loadAllCustomers()
        },
        error: function (ob, txtStatus, error) {
            console.log(error);
            console.log(txtStatus);
            console.log(ob);
        }
    });
});

//AllCustomers
function loadAllCustomers() {


    console.log("working load all customer fuction")

    $('#tblCustomer').empty();

    $.ajax({
        method: "GET",
        url: "http://localhost:8000/api/customer",
        // async: true,
        success: function (res) {
            let data = res.data;
            console.log(data)
            console.log("Response check")
            for (var i in res){
                let id = res[i].id;
                let name = res[i].name;
                let address = res[i].address;
                let nic = res[i].nic;
                let contact = res[i].contact;

                var row=`<tr> <td>${id}</td> <td>${name}</td><td>${address}</td><td>${nic}</td><td>${contact}</td></tr>`;
                $('#tblCustomer').append(row);

            }
        },
        error: function (ob, txtStatus, error) {
            console.log(error);
            console.log(txtStatus);
            console.log(ob);
        }
    })
}

// search customer
$("#btnSearch").click(function () {

    console.log("work btnSearchCar");

    let custID = $("#custID").val();
    $.ajax({
        method: "GET",
        url: "http://localhost:8000/api/customer/" + custID,
        success: function (res) {
            console.log(res);
                let c = res.data;
                // set details to input fields
                //$("#custID").val(res.data.);
                $("#custName").val(res.name);
                $("#address").val(res.address);
                $("#nic").val(res.nic);
                $("#contact").val(res.contact);
        },
        error: function (ob, txtStatus, error) {
            console.log(error);
            console.log(txtStatus);
            console.log(ob);
        }
    });
});
