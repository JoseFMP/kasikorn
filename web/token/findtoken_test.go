package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFindTokenId(t *testing.T) {

	tokenId := FindTokenId(preloginPayload)
	assert.NotNil(t, tokenId)
	assert.Equal(t, "1515123653", *tokenId)
}

const preloginPayload = `





























<!DOCTYPE html>
<html lang="th" id="loginContents">
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
        <meta http-equiv="Pragma" content="no-cache" />
        <meta http-equiv="Cache-Control" content="no-cache, no-store" />
        <meta http-equiv="Expires" content="0" />
        <meta http-equiv="X-UA-Compatible" content="IE=edge" />
        <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1" />  
        
        
        
                <title>K-Cyber for SME Login</title>
        
        <script type="text/javascript">
        var locale = "en";    
        if(isIE()&&isIE()<10) {
            if(locale === "en"){
                alert("Support Internet Explorer Version 10 and above.");
            }else{
                alert("รองรับ Internet Explorer เวอร์ชั่น 10 หรือมากกว่า");                
            }    
            
        }
        function isIE () {
          var myNav = navigator.userAgent.toLowerCase();
          return (myNav.indexOf('msie') != -1) ? parseInt(myNav.split('msie')[1]) : false;
        }
        </script>  
        <link rel="shortcut icon" href="/K-Online/images/favicon/kbank.ico"/>
        <link rel="apple-touch-icon" href="/K-Online/images/favicon/icon_iPad.png"/>
        
        <link rel="stylesheet" href="/K-Online/css/redesign/bootstrap.css">
        <link rel="stylesheet" href="/K-Online/css/redesign/fonts.css">
        <link rel="stylesheet" href="/K-Online/css/redesign/font-icokbank.css">
        <link rel="stylesheet" href="/K-Online/css/redesign/perfect-scrollbar.min.css">
        <link rel="stylesheet" href="/K-Online/css/redesign/jquery.mmenu.all.min.css">
        <link rel="stylesheet" href="/K-Online/css/redesign/masterslider.min.css">
        <link rel="stylesheet" href="/K-Online/css/redesign/owl.carousel.css">
        <link rel="stylesheet" href="/K-Online/css/redesign/slick.css">
        <link rel="stylesheet" href="/K-Online/css/redesign/magnific-popup.css">
        <link rel="stylesheet" href="/K-Online/css/redesign/chosen-custom.css">       
        <link rel="stylesheet" href="/K-Online/css/redesign/theme.css">
        <link rel="stylesheet" href="/K-Online/css/redesign/theme-rwd.css">  
	<link rel="stylesheet" href="/K-Online/css/redesign/k-online.css">		
        <link rel="stylesheet" href="/K-Online/css/new_login/jcarousel.basic.css">
        
        <script type="text/javascript" src="/K-Online/scripts/redesign/modernizr.min.js"></script>
        <script type="text/javascript" src="/K-Online/scripts/jquery.js"></script>
        <script type="text/javascript" src="/K-Online/scripts/redesign/jquery.easing.1.3.min.js"></script>
        <script type="text/javascript" src="/K-Online/scripts/validator.js?v=1515123653"></script>
        <script type="text/javascript" src="/K-Online/scripts/CapsLock.js?v=1515123653"></script>
        <script type="text/javascript" src="/K-Online/scripts/pa-libs/pa-tools.js?v=1515123653"></script>
        <script type="text/javascript" src="/K-Online/scripts/pa-libs/pa-libs.js?v=1515123653"></script>
        
        <script type="text/javascript" src="/K-Online/scripts/redesign/fixto.min.js"></script>
        <script type="text/javascript" src="/K-Online/scripts/redesign/perfect-scrollbar.jquery.min.js"></script>
        <script type="text/javascript" src="/K-Online/scripts/redesign/masterslider.min.js"></script>
        <script type="text/javascript" src="/K-Online/scripts/redesign/chosen.jquery.min.js"></script>
        <script type="text/javascript" src="/K-Online/scripts/redesign/jquery.magnific-popup.min.js"></script>
        <script type="text/javascript" src="/K-Online/scripts/redesign/global.js"></script>
        <script type="text/javascript" src="/K-Online/scripts/redesign/k-online.js"></script>
        
        <script type="text/javascript" src="/K-Online/scripts/libs/jquery.jcarousel.min.js"></script>
        <script type="text/javascript" src="/K-Online/scripts/libs/jquery.jcarousel-autoscroll.min.js"></script>
        <script type="text/javascript" src="/K-Online/scripts/libs/jcarousel.basic.js"></script>
        <script type="text/javascript" src="/K-Online/scripts/libs/jquery.zrssfeed.js"></script>

        <script type="text/javascript">
            $(document).keypress(function(e) {
                if (e.keyCode == 13) {
                    doAction('authenticate');
                }
            });
            
            CapsLock.addListener(function(status) {
                document.getElementById('capsMsg').style.display = (status ? 'block' : 'none');
            });
            
            $(document).ready(function(){
                $('#userName').prop('autocomplete', 'off');
                $('#password').prop('autocomplete', 'off');
                $('#userName').focus();
                $("form").attr("autocomplete", "off");
                
                $('#loginBtn').click(function(){
                    doAction('authenticate');
                });
                
                $('.open-popup-contact').click(function() {
                    $('#form-inform').stop().css('opacity',1).show();
                    $('#inform-message').stop().css('opactiy',0).hide();
                });
                
                $('#inform-message').hide();
                $('#form-inform').submit(function(e) {
                    if(validateTrojanReport()){
                        e.preventDefault();
                        sendMail(this); 
                    }
                    return false;
                });
                
                $('#firstname').keyup(function(){
                    if ($(this).val().length > 0) {
                        $(this).css('border', '1px solid #eee');
                        $('#errorFirstname').html('');
                    }
                });
                $('#lastname').keyup(function(){
                    if ($(this).val().length > 0) {
                        $(this).css('border', '1px solid #eee');
                        $('#errorLastname').html('');
                    }
                });
                $('#email').keyup(function(){
                    if ($(this).val().length > 0) {
                        $(this).css('border', '1px solid #eee');
                        $('#errorEmail').html(''); 
                    }
                });
                $('#mobileNo').keyup(function(){
                    if ($(this).val().length > 0) {
                        $(this).css('border', '1px solid #eee');
                        $('#errorMobileno').html('');
                    }
                });
                $('#msg').keyup(function(){
                    if ($(this).val().length > 0) {
                        $(this).css('border', '1px solid #eee');
                        $('#errorMessage').html('');
                    }
                });
                var locale = $("#locale").val();
                if (locale == null || locale == "") {
                    $('#locale').val("th");
                    locale = "th";
                }
                
                if (locale == "th") {
                    $('#localeTh').css('cursor', 'default');
                    $('#localeTh').css('text-decoration', 'underline');
                    $('#localeEn').css('cursor', 'pointer');
                } else {
                    $('#localeEn').css('cursor', 'default');
                    $('#localeEn').css('text-decoration', 'underline');
                    $('#localeTh').css('cursor', 'pointer');
                }
                
                $('span.btn_locale').click(function(){
                    var localeId = this.id;
                    var styleBtnEn = $('#localeEn').css('cursor');
                    var styleBtnTh = $('#localeTh').css('cursor');
                    if (localeId.indexOf('En') != -1 && styleBtnEn != 'default') {
                        $('#locale').val("en");
                        $(this).css('cursor', 'default');
                        $('#localeTh').css('cursor', 'pointer');
                        $('#cmd').val('lang');
                        doChangeLocale($("#locale").val(),$('#custType').val());
                    } else if (localeId.indexOf('Th') != -1 && styleBtnTh != 'default') {
                        $('#locale').val("th");
                        $(this).css('cursor', 'default');
                        $('#localeEn').css('cursor', 'pointer');
                        $('#cmd').val('lang');
                        doChangeLocale($("#locale").val(),$('#custType').val());
                    }
                });
                
                $('.keypress_number').keypress(function(event){
                    if (event.which != null || (event.charCode != null || event.keyCode != null)) {
                        event.which = event.charCode != null ? event.charCode : event.keyCode;
                        if (event.which >= 48 && event.which<=57 || event.which == 8 || event.which == 0) {
                            return true;
                        }
                    }
                    return false;
                });
                
                pajsfile("/K-Online", "footer.js?v=1515123653", "js");
                
                $('#menuOther').click(function() {
                   $('#firstname').val('');
                   $('#firstname').css('border', '1px solid #eee');
                   $('#errorFirstname').html('');
                   
                   $('#lastname').val('');
                   $('#lastname').css('border', '1px solid #eee');
                   $('#errorLastname').html('');
                   
                   $('#email').val('');
                   $('#email').css('border', '1px solid #eee');
                   $('#errorEmail').html('');
                   
                   $('#mobileNo').val('');
                   $('#mobileNo').css('border', '1px solid #eee');
                   $('#errorMobileno').html('');
                   
                   $('#msg').val('');
                   $('#msg').css('border', '1px solid #eee');
                   $('#errorMessage').html('');
                });
            });
            
            function doAction (cmd) {
                if(validate()){
                    $('#cmd').val(cmd);
                    $('form').submit();
                }
            }

            function validate() {                
                var validator = new Validator('en');                
                validator.setErrorPanel($('#errorText')[0]);
                validator.setSuccessPanel($('#errorText')[0]);
                isNull($('#userName').val(), $('#password').val());

                validator.addField($('#userName')[0], 'User ID', { required: true }  );
                validator.addField($('#password')[0], 'Password', { required: true }  );
                validator.validateField();
                if (!validator.getDelegator().valid) {                    
                    validator.displayError();
                    return false;
                }

                if('' == 'true') {
                    validator.addField($('#captcha')[0], 'Characters', { required: true}  );
                    validator.validateField();
                    if (!validator.getDelegator().valid) {                          
                        validator.displayError();                        
                        return false;
                    }
                }

                $.post("analytic.jsp", { cmd: 'c', paObjects: encode64(encryptedElement()) });

                return true;
            }
            
            function validateTrojanReport(){
            $('#errorFirstname').html('');
            $('#errorLastname').html('');
            $('#errorEmail').html('');
            $('#errorMobileno').html('');
            $('#errorMessage').html('');

            var isSuccess = true;
            var msgRequireField = 'Please fill your information';
            var msgEmailFormat = 'Invalid e-mail address';
            var msgMobileNoFormat = 'Invalid Mobile number';
                if($('#msg').val()===''){
                    var err = $('#errorMessage')[0];
                    err.innerHTML = msgRequireField; 
                    $('#msg').css('border', '1px solid #f00');
                    isSuccess = false;
                }  
                
                if($('#mobileNo').val()===''){
                    var err = $('#errorMobileno')[0];
                    err.innerHTML = msgRequireField;    
                    $('#mobileNo').css('border', '1px solid #f00');
                    isSuccess = false;
                }else if($('#mobileNo').val().trim().length!==10){
                    var err = $('#errorMobileno')[0];
                    err.innerHTML = msgMobileNoFormat;
                    $('#mobileNo').css('border', '1px solid #f00');
                    isSuccess = false;
                } 
                
                if($('#email').val()===''){
                    var err = $('#errorEmail')[0];
                    err.innerHTML = msgRequireField;
                    $('#email').css('border', '1px solid #f00');
                    isSuccess = false;
                }else if(validateEmail($('#email').val().trim())){
                    var err = $('#errorEmail')[0];
                    err.innerHTML = msgEmailFormat;
                    $('#email').css('border', '1px solid #f00');
                    isSuccess = false;
                }
                
                if($('#lastname').val()===''){
                    var err = $('#errorLastname')[0];
                    err.innerHTML = msgRequireField;
                    $('#lastname').css('border', '1px solid #f00');
                    isSuccess = false;
                }
                
                if($('#firstname').val()===''){
                    var err = $('#errorFirstname')[0];
                    err.innerHTML = msgRequireField;
                    $('#firstname').css('border', '1px solid #f00');
                    isSuccess = false;
                }                   
                                
                return isSuccess;
            }
            
            function validateEmail(value){
                var emailRegEx = /^[A-Z0-9_-]+(?:\.[A-Z0-9_-]+)*@(?:[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?\.)+[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])$/i;
                return !isValidEmail(value,emailRegEx);
            }
            
            function isValidEmail(value, regexp) {
                if (value != "" && value != null) {
                    var myRe = regexp.exec(value);
                    if (!myRe) {
                        return false;
                    } else {
                        return true;
                    }
                }
                return true;
            }
            
            function sendMail(e){
                $.post({
                    type: 'POST',
                    url: 'trojanReport.do',
                    data: $("#form-inform").serialize(),
                    cache: false,
                    success: function(resp){
                        $(e).fadeTo(320, 0, function(){
                            if ("success" === resp.trim()) {
                            } else {                                
                                $('#msgThx').html('');
                                $('#result-message').html('Sorry, unable to process at this moment, please try again later or contact K-Contact Center +662-8888888');
                            }
                            $(e).hide();
                            $('#inform-message').css({
                                opacity: 0
                            }).show().fadeTo(560, 1);
                        });
                    }
                });
            }
            
            function doChangeLocale (locale,custType) {
                window.location = "/K-Online/login.jsp?lang=" + locale + "&type=" + custType;
            }

        </script>

		<style type="text/css">
            <!--
            body {
                overflow:hidden;
            }
            -->
        </style>

    </head>
    
    <body>
        <div id="reward-feed" style="display:none"></div>
        <div id="news-feed" style="display:none"></div>
        
        
            <div id="smart-app-ios" class="smart-app">
                <a href="#" class="a-close smart-app-close"><i class="ic ic-close"></i></a>
                <img src="/K-Online/images/redesign/KPlusSME.png" alt="" width="63">
                <span class="smart-app-name">K PLUS SME<br><small>KASIKORNBANK PCL.</small></span>
                <a href="https://appsto.re/th/4myCgb.i" target="_blank" class="a-view">View</a>
            </div>

            <div id="smart-app-android" class="smart-app">
                <a href="#" class="a-close smart-app-close"><i class="ic ic-close"></i></a>
                <img src="/K-Online/images/redesign/KPlusSME.png" alt="" width="63">
                <span class="smart-app-name">K PLUS SME<br><small>KASIKORNBANK PCL.</small></span>
                <a href="https://play.google.com/store/apps/details?id=com.kasikorn.sme.mbanking&hl=en" target="_blank" class="a-view">View</a>
            </div>
        
        
        <div id="page" class="icw login-page">
            <div id="menu_line" style="display:none" class="line_612f6f3ac31b89a84e80ebc48ee05cb09d6c9948e92ee8bf669b06faed303401"></div>
            <header id="header" class="header hide line_c2NyaXB0LGZvcm0saW5wdXQsYSxpbWcsaWZyYW1lLGhpZGRlbixvYmplY3Q="></header>
            <div id="login-content" class="login-content line_d3JpdGV8Y3JlYXRlRWxlbWVudA==">
                <div id="login-container" class="login-container line_I3xLLU9ubGluZXxzZWN1cmUua2FzaWtvcm5iYW5rZ3JvdXAuY29tfHd3dy5rYXNpa29ybmJhbmsu
Y29tfHd3dy5nb29nbGUtYW5hbHl0aWNzLmNvbQ==">                    
                    <div class="login-inner">
                        <div id="navbar" class="navbar line_aW1nLHNjcmlwdCxpZnJhbWUsYQ==">
                            <span id="localeTh" name="localeTh" class="lang btn_locale" alt="TH" border="0" style="font-size: 19px">ไทย</span>
                            <span class="btn_locale" border="0">|</span>
                            <span id="localeEn" name="localeEn" class="lang btn_locale" alt="EN" border="0" style="font-size: 16px">ENG</span>
                            <a data-mfp-src="#popup-content" class="menu open-popup-content" id="menuOther" name="menuOther"><i class="ic"></i></a>
                        </div>

                        <div class="logo">
                            <a href="http://www.kasikornbank.com/en">
                                <img src="/K-Online/images/redesign/kasikornbank.png" width="205" alt="KASIKORNBANK" class="img-logo" />
                            </a>
                        </div>

                        <div class="form-box form-box-sm">
                            <div class="heading">
                                
                                
                                    K-Cyber for SME
                                
                            </div>
                            <div class="entryform entrycontent">
                                <div id="action" class="line_YWpheDo0LGV2YWw6MiwucGhwOjE=" style="display: none"></div>
                                <form name="loginActionForm" method="post" action="/K-Online/login.do" class="icw-form">
                                    <input type="hidden" name="tokenId" id="tokenId" value="1515123653"/>
                                    <div class ="errorMsg" id="errorText" style="margin-bottom: 10px; font-size: 14px; color: #f00;"></div>
                                    <div class="icw-input">
                                        <input type="text" name="userName" maxlength="20" value="" id="userName">
                                        <label class="icw-label overlabel">User ID</label>
                                    </div>
                                    <div class="icw-input">
                                        <input type="password" name="password" maxlength="30" value="" id="password">
                                        <label class="icw-label overlabel">Password</label>                                        
                                    </div>
                                    <div id="capsMsg" style="color: #800000; display: none" width="100%">
                                        <img src="/K-Online/images/notify-icon.png" width="12" height="12"/>&nbsp;CAPS LOCK is on!
                                    </div>
                                    
                                    <div class="action-form">
                                        <button type="button" id="loginBtn" class="btn btn-block">
                                            
                                            Log In
                                        </button>
                                        <div class="other">
                                            
                                            
                                                <a href="/K-Online/onlineUnlock.do?cmd=init&type=sme">
                                                    
                                                    
                                                        Unlock User ID
                                                    
                                                </a>
                                            
                                        </div>
                                    </div>
                                    <input type="hidden" name="cmd" value="authenticate" id="cmd">
                                    <input type="hidden" name="locale" value="en" id="locale">
                                    <input type="hidden" name="custType" value="sme" id="custType">
                                    <input type="hidden" name="app" value="0">
                                </form>
                            </div>
                        </div>

                        <div class="action-form">
                            <div class="other">
                                <ul class="list-style-none list-inline">
                                    <li>
                                        
                                        Not have User ID?
                                    </li>
                                    <li>
                                        
                                        
                                            <a href="/K-Online/smeRegistration/firstPage.jsp" class="btn btn-outline">    
                                            
                                                
                                                Services Registration
                                            </a>
                                    </li>
                                </ul>
                            </div>
                        </div>

                        <div class="icw-notice">
                            <p>
                                
                                
                                    KBank does not have a policy to send SMS/MMS/Email embedded link.
                                
                            </p>
                        </div>
                    </div>
                    <div class="push"></div>
                </div>

                <div id="slogan" class="slogan line_dG9rZW5JZHx1c2VyTmFtZXxwYXNzd29yZHxjbWR8bG9jYWxlfGN1c3RUeXBlfGFwcHxjYXB0Y2hh
fGZpcnN0bmFtZXxsYXN0bmFtZXxlbWFpbHxtb2JpbGVOb3xzdWJtaXRCdG4=">
                    <img src="/K-Online/images/redesign/slogan.png" class="img">
                </div>
            </div>

            <div id="preview-slider" class="preview-slider">
                <div class="master-slider">
                    
                    
                        <div class="ms-slide" data-delay="6">
                            <img src="/K-Online/images/redesign/blank.gif" data-src="https://www.kasikornbank.com/th/personal/Digital-banking/PublishingImages/KCB/K-Cyber_SME_preview.jpg"/>
                            <div class="ms-layer" style="top:0; left:0;" data-duration="2500" data-effect="bottom(40)" data-ease="easeOutExpo">
                                <div class="entrycontent">
                                    <div class="h1 heading futuraHeavyBT c-white">K-Cyber for SME</div>
                                    <div class="h2 sub-heading c-white">
                                        
                                        
                                            Internet Banking Service for juristic persons or business owners
                                        
                                    </div>
                                </div>
                            </div>
                        </div>
                    
                </div>
            </div>

            <div class="clr"></div>

            <div id="popup-content" class="icw mfp-hide">
                <div class="main-container">
                    <div class="container">
                        <div class="two-columns tabs-group">
                            <div id="" class="sidebar sidebar-fixed">
                                <div class="sidebar-inner">
                                    <h3 class="heading">
                                        
                                        
                                            K-Cyber for SME
                                        
                                    </h3>
                                    <ul class="nav tabs-list">
                                        <li class="current">
                                            <a href="#security-tips" class="tab-trigger">
                                                
                                                
                                                    Security Tips
                                                
                                            </a>
                                        </li>
                                        <li>
                                            <a href="#qr-otp" class="tab-trigger">
                                                
                                                
                                                    QR-OTP Set up Instruction
                                                
                                            </a>
                                        </li>    
                                        <li>
                                            <a href="#inform" class="tab-trigger">
                                                
                                                
                                                    Report Phishing Email & Website
                                                
                                            </a>
                                        </li>
                                    </ul>
                                </div>
                            </div>

                            <div class="column-main">
                                <div class="entrycontent">
                                    <div id="security-tips" class="tab-content tab-content-current">
                                        
                                        
                                            <div class="content-intro">
                                                <div class="h2 heading c-green">Security Tips</div>
                                                <h3 class="sub-heading">Security Tips for K-Cyber and K-Cyber for SME</h3>
                                            </div>
                                            <div class="a-center">
                                                <img src="https://www.kasikornbank.com/th/personal/Digital-banking/PublishingImages/KCB/K-Cyber_SecurityTip_EN.jpg" />
                                            </div>
                                        
                                    </div>

                                    <div id="qr-otp" class="tab-content">
                                        
                                        
                                            <div class="content-intro">
                                                <div class="h2 heading c-green">QR-OTP Set up Instruction</div>
                                            </div>
                                            <p>QR-OTP, the new security technology developed for K-Cyber and K-Cyber for SME services to confirm important transactions. See QR-OTP set up instruction below:</p>
                                            <br>
                                            <div class="a-center">
                                                <img src="https://www.kasikornbank.com/th/personal/Digital-banking/PublishingImages/KCB/K-Secure_EN.jpg" />
                                            </div>
                                            <br>
                                        
                                    </div>
                                    <div id="inform" class="tab-content">
                                        <div class="content-intro">
                                            <div class="h2 heading c-green">
                                                
                                                
                                                    Report Phishing Email & Website
                                                
                                            </div>
                                        </div>

                                        <form name="emailTrojanReportForm" id="form-inform" method="post" action="/K-Online/trojanReport.do" class="icw-form">
                                            <ul class="icw-fields">
                                                <li class="icw-field">
                                                    <div class="icw-input">                                                        
                                                        <input type="text" name="firstname" value="" id="firstname">
                                                        <label  class="icw-label overlabel">
                                                            
                                                            
                                                                First name
                                                            
                                                        </label>
                                                        <div id="errorFirstname" style="margin-bottom: 10px; font-size: 14px; color: red;"></div>
                                                    </div>
                                                </li>

                                                <li class="icw-field">
                                                    <div class="icw-input">                                                        
                                                        <input type="text" name="lastname" value="" id="lastname">
                                                        <label  class="icw-label overlabel">
                                                            
                                                            
                                                                Last name
                                                                                                                        
                                                        </label>
                                                        <div id="errorLastname" style="margin-bottom: 10px; font-size: 14px; color: red;"></div>
                                                    </div>
                                                </li>

                                                <li class="icw-field">
                                                    <div class="icw-input">
                                                        <input type="text" name="email" value="" id="email">
                                                        <label  class="icw-label overlabel">
                                                            
                                                            
                                                                Email
                                                            
                                                        </label>
                                                        <div id="errorEmail" style="margin-bottom: 10px; font-size: 14px; color: red;"></div>
                                                    </div>
                                                </li>

                                                <li class="icw-field">
                                                    <div class="icw-input">
                                                        <input type="text" name="mobileNo" value="" id="mobileNo" class="keypress_number">
                                                        <label  class="icw-label overlabel">
                                                            
                                                            
                                                                Phone number
                                                            
                                                        </label>
                                                        <div id="errorMobileno" style="margin-bottom: 10px; font-size: 14px; color: red;"></div>
                                                    </div>
                                                </li>

                                                <li class="icw-field icw-field-full">
                                                    <div class="icw-input">
                                                        <textarea name="msg" id="msg"></textarea>
                                                        <label  class="icw-label overlabel">
                                                            
                                                            
                                                                Message
                                                            
                                                        </label>
                                                        <div id="errorMessage" style="margin-bottom: 10px; font-size: 14px; color: red;"></div>
                                                    </div>
                                                </li>
                                            </ul>

                                            <div class="action-form">
                                                
                                                
                                                    <input type="submit" name="submitBtn" class="btn btn-md" value="Send message">
                                                
                                            </div>         
                                        </form>

                                        <div id="inform-message" class="form-message a-center">
                                            
                                            
                                                <div id="msgThx" class="h3">Thank you for your information</div>
                                                <br>
                                                <p id="result-message">We have received your information.<br>Any enquiries, please call &nbsp; <a href="tel:028888888" class="link-icon"><i class="ic ic-phone"></i>02-8888888</a> or &nbsp; <a href="mailto:info@kasikornbank.com" class="link-icon"><i class="ic ic-envelope"></i>info@kasikornbank.com</a></p>
                                                
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <footer class="footer hide"></footer>
        </div>               
        <script> 
    (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){ 
    (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o), 
    m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m) 
    })(window,document,'script','//www.google-analytics.com/analytics.js','ga'); 

    ga('create', 'UA-10493652-15', 'auto'); 
    ga('send', 'pageview');
</script>
    </body>
</html>

`
