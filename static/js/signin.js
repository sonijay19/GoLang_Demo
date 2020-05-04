function docin() {
	console.log('in js');
    $.ajax({
        method: 'post',
        url: '/SignUp/',
        data:({
        	drName : $('#drName').val(),
        	drmobNo : $('#drmobNo').val(),
        	drEmail : $('#drEmail').val(),
        	drPassword : $('#drPassword').val(),
        	drLicense : $('#drLicense').val(),
        	drCity : $('#drCity').val(),
        	drPincode : $('#drPincode').val(),
        	domain : $('#domain_doc_sign').val()
        }),
        success: function(responseText) {
            if(responseText == "1"){
            	window.location.assign("/dashboard/");
            }
            else{
            	$("#jay_doc_sign").html(responseText);
            }
        }
    });
    return false;
};



function chemin() {
	console.log('in js');
    $.ajax({
        method: 'post',
        url: 'SignUp',
        data:({
        	cName : $('#cName').val(),
        	cmobNo : $('#cmobNo').val(),
        	cEmail : $('#cEmail').val(),
        	cAddress : $('#cAddress').val(),
        	cCity : $('#cCity').val(),
        	cPincode : $('#cPincode').val(),
        	cPassword : $('#cPassword').val(),
        	domain : $('#domain_chem_sign').val()
        }),
        success: function(responseText) {
            if(responseText == "1"){
            	window.location.assign("/dashboard/");
            }
            else{
            	$("#jay_chem_sign").html(responseText);
            }
        }
    });
    return false;
};


function labin() {
	console.log('in js');
    $.ajax({
        method: 'post',
        url: 'SignUp',
        data:({
        	lName : $('#lName').val(),
        	lmobNo : $('#lmobNo').val(),
        	lEmail : $('#lEmail').val(),
        	lLicense : $('#lLicense').val(),
        	lAddress : $('#lAddress').val(),
        	lCity : $('#lCity').val(),
        	lPincode : $('#lPincode').val(),
        	lPassword : $('#lPassword').val(),
        	domain : $('#domain_lab_sign').val()
        }),
        success: function(responseText) {
            if(responseText == "1"){
            	window.location.assign("helth/Doctor/dashboard.jsp");
            }
            else{
            	$("#jay_lab_sign").html(responseText);
            }
        }
    });
    return false;
};





function doclogin1() {
	console.log('in js');
	console.log($('#drEmail_log').val())
	console.log($('#drPassWord_log').val())
    $.ajax({
        method: 'post',
        url: '/LogIn/',
        data:({
        	drEmail : $('#drEmail_log').val(),
        	drPassWord : $('#drPassWord_log').val(),
        	domain : $('#domain_log_dr').val()
        }),
        success: function(responseText) {
            if(responseText == "1"){
            	window.location.assign("/dashboard/");
            }
            else{
            	$("#jay_log_dr").html(responseText);
            }
        }
    });
    return false;
};


function chemlogin() {
	console.log('in js');
    $.ajax({
        method: 'post',
        url: 'SignIn',
        data:({
        	cEmail : $('#cEmail_log').val(),
        	cPassword : $('#cPassword_log').val(),
        	domain : $('#domain_log_chem').val()
        }),
        success: function(responseText) {
            if(responseText == "1"){
            	window.location.assign("helth/Doctor/dashboard.jsp");
            }
            else{
            	$("#jay_log_chem").html(responseText);
            }
        }
    });
    return false;
};


function lablogin() {
	console.log('in js');
    $.ajax({
        method: 'post',
        url: 'SignIn',
        data:({
        	lEmail : $('#log_lEmail').val(),
        	lPassword : $('#log_lPassword').val(),
        	domain : $('#domain_log_lab').val()
        }),
        success: function(responseText) {
            if(responseText == "1"){
            	window.location.assign("helth/Doctor/dashboard.jsp");
            }
            else{
            	$("#jay_log_lab").html(responseText);
            }
        }
    });
    return false;
};
