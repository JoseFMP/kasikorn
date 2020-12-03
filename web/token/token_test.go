package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanFindToken(t *testing.T) {

	token := FindToken(tokenTestData)
	assert.NotNil(t, token)
	assert.Equal(t, "0.21369009104596925", *token)
}

const tokenTestData = `




<style type='text/css'> body { margin-top:0px; margin-bottom:0px; margin-right:0px; margin-left:16px; overflow-y:scroll; }</style>









<html>
<head>
<!-- #BeginEditable "doctitle" --> 
<title>
Welcome to K-Cyber Banking
</title>
<!-- #EndEditable -->
<meta http-equiv="Content-Type" content="text/html; charset=windows-874">

<meta http-equiv="Expires" content="now">
<meta http-equiv="Pragma" content="no-cache">
<meta http-equiv="Pragma-directive" content="no-cache">
<meta http-equiv="cache-directive" content="no-cache">

<style type="text/css">
</style>
<link rel="stylesheet" href="/retailstatic/css/stylesheet.css" type="text/css">
<LINK rel="stylesheet" href="/retailstatic/css/qstyle.css" rel=stylesheet type=text/css>

    




<Script language = "JavaScript">
function openWebChatWindow(){	
	var gLanguage = "english";
	if((document.frmWebChat.custData.value == "") && (gLanguage=="english")){
		window.open("/retail/notification/webchat_error.jsp","WebChatError","width=780,height=450,top=0,left=0,toolbar=no,location=no,directories=no,status=no,menubar=no,scrollbars=yes,resizable=yes");
	}else if((document.frmWebChat.custData.value == "") && (gLanguage=="thai")){	
		window.open("/retail/notification/webchat_error_th.jsp","WebChatError","width=780,height=450,top=0,left=0,toolbar=no,location=no,directories=no,status=no,menubar=no,scrollbars=yes,resizable=yes");
	}else{
		document.frmWebChat.submit();
	}
}
</Script>

<FORM NAME="frmWebChat" ACTION="https://rt04.kasikornbank.com/website/servlet/webchatinit" TARGET="_blank" METHOD="post">
<INPUT type="hidden" name="custData" value="bgu6SRKZ0QzT2Kkcj9lWvStkej41qdDZK/qawETl0JQY/vLahG+7sTULoYXHRaUCiuSNAeTCDTJJq+E18IgeJ4WEZIoTf9S8CfQfIpgicmTH2jzXkyWB26+LqMsp18nBV1yZynJ24MpPq/FTmm7eFviQyLypKBA9">
</FORM>


<script language="javascript" src="/retailstatic/jslib/slide_menu.js"></script>
<script language="javascript" src="/retailstatic/jslib/left_menu.js"></script>
<script language="javascript" src="/retailstatic/jslib/menu.js"></script>
<script language="javascript" src="/retailstatic/jslib/top_menu.js"></script>
<script language="javascript" src="/retailstatic/jslib/mid_menu.js"></script> 
<script language="javascript" src="/retailstatic/jslib/buttom_menu.js"></script>
<script language="javascript" src="/retailstatic/jslib/tfb.js?version=20170322"></script>
<script language="javascript" src="/retailstatic/jslib/calendar.js"></script>
<script language="javascript" src="/retailstatic/jslib/constant.js"></script>




<script language="JavaScript">
var printURL;
function MM_goToURL() { //v3.0
  var i, args=MM_goToURL.arguments; document.MM_returnValue = false;
  for (i=0; i<(args.length-1); i+=2) eval(args[i]+".location='"+args[i+1]+"'");
}
function MM_jumpMenu(targ,selObj,restore){ //v3.0
  eval(targ+".location='"+selObj.options[selObj.selectedIndex].value+"'");
  if (restore) selObj.selectedIndex=0;
}

function MM_openBrWindow(theURL,winName,features) { //v2.0
  window.open(theURL,winName,features);
}

function MM_preloadImages() { //v3.0
  var d=document; if(d.images){ if(!d.MM_p) d.MM_p=new Array();
    var i,j=d.MM_p.length,a=MM_preloadImages.arguments; for(i=0; i<a.length; i++)
    if (a[i].indexOf("#")!=0){ d.MM_p[j]=new Image; d.MM_p[j++].src=a[i];}}
}

function MM_swapImgRestore() { //v3.0
  var i,x,a=document.MM_sr; for(i=0;a&&i<a.length&&(x=a[i])&&x.oSrc;i++) x.src=x.oSrc;
}

function MM_findObj(n, d) { //v4.01
  var p,i,x;  if(!d) d=document; if((p=n.indexOf("?"))>0&&parent.frames.length) {
    d=parent.frames[n.substring(p+1)].document; n=n.substring(0,p);}
  if(!(x=d[n])&&d.all) x=d.all[n]; for (i=0;!x&&i<d.forms.length;i++) x=d.forms[i][n];
  for(i=0;!x&&d.layers&&i<d.layers.length;i++) x=MM_findObj(n,d.layers[i].document);
  if(!x && d.getElementById) x=d.getElementById(n); return x;
}

function MM_swapImage() { //v3.0
  var i,j=0,x,a=MM_swapImage.arguments; document.MM_sr=new Array; for(i=0;i<(a.length-2);i+=3)
   if ((x=MM_findObj(a[i]))!=null){document.MM_sr[j++]=x; if(!x.oSrc) x.oSrc=x.src; x.src=a[i+2];}
}

function openPrintWindow()
{


window.open(printURL,"Print","width=780,height=450,top=0,left=0,toolbar=no,location=no,directories=no,status=no,menubar=no,scrollbars=yes,resizable=yes");

}

function nothing(){
}
function toggleBullet(elm) {
 var newDisplay = "none";
 var e = elm.nextSibling; 
 while (e != null) {
  if (e.tagName == "P" || e.tagName == "p") {
   if (e.style.display == "none") newDisplay = "block";
   break;
  }
  e = e.nextSibling;
 }
 while (e != null) {
  if (e.tagName == "P" || e.tagName == "p") e.style.display = newDisplay;
  e = e.nextSibling;
 }
}

function toggleBulletUL(elm) {
 var newDisplay = "none";
 var e = elm.nextSibling; 
 while (e != null) {
  if (e.tagName == "UL" || e.tagName == "ul") {
   if (e.style.display == "none") newDisplay = "block";
   break;
  }
  e = e.nextSibling;
 }
 while (e != null) {
  if (e.tagName == "UL" || e.tagName == "ul") e.style.display = newDisplay;
  e = e.nextSibling;
 }
}


function collapseAll() {
  var lists = document.getElementsByTagName('P');
  for (var j = 0; j < lists.length; j++) 
   lists[j].style.display = "none";
  lists = document.getElementsByTagName('p');
  for (var j = 0; j < lists.length; j++) 
   lists[j].style.display = "none";
  var e = document.getElementById("root");
  e.style.display = "block";
}

function collapseAllUL() {
  var lists = document.getElementsByTagName('UL');
  for (var j = 0; j < lists.length; j++) 
   lists[j].style.display = "none";
  lists = document.getElementsByTagName('ul');
  for (var j = 0; j < lists.length; j++) 
   lists[j].style.display = "none";
  var e = document.getElementById("root");
  e.style.display = "block";
}

var caution = false
function setCookie(name, value, expires, path, domain, secure) {
	var curCookie = name + "=" + escape(value) +
		((expires) ? "; expires=" + expires.toGMTString() : "") +
		((domain) ? "; domain=" + domain : "") +
		("; path=/") +
		((secure) ? "; secure" : "")
	if (!caution || (name + "=" + escape(value)).length <= 4000)
		document.cookie = curCookie
	else
		if (confirm("Cookie exceeds 4KB and will be cut!"))
			document.cookie = curCookie
}
function getCookie(name) {
	var prefix = name + "="
	var cookieStartIndex = document.cookie.indexOf(prefix)
	if (cookieStartIndex == -1)
		return null
	var cookieEndIndex = document.cookie.indexOf(";", cookieStartIndex + prefix.length)
	if (cookieEndIndex == -1)
		cookieEndIndex = document.cookie.length
	return unescape(document.cookie.substring(cookieStartIndex + prefix.length, cookieEndIndex))
}
function deleteCookie(name, path, domain) {
	if (getCookie(name)) {
		document.cookie = name + "=" +
		((domain) ? "; domain=" + domain : "") +
		"; expires=Thu, 01-Jan-70 00:00:01 GMT"
	}
}
function fixDate(date) {
	var base = new Date(0)
	var skew = base.getTime()
	if (skew > 0)
		date.setTime(date.getTime() - skew)
}
var storedValue = getCookie("outline");
function initimages() {
		standard = new MakeArray(11);
		over = new MakeArray(11);
	
		standard[1].src="/retailstatic/images/service_menu/accman_en.gif";
		standard[2].src="/retailstatic/images/service_menu/cc_en.gif";
		standard[3].src="/retailstatic/images/service_menu/ft_en.gif";
		standard[4].src="/retailstatic/images/service_menu/bp_en.gif";
		standard[5].src="/retailstatic/images/service_menu/cheque_en.gif";
		standard[6].src="/retailstatic/images/service_menu/ew_en.gif";
		standard[7].src="/retailstatic/images/service_menu/mp_en.gif";
	//	standard[8].src="/retailstatic/images/service_menu/ln_en.gif";
		standard[10].src="/retailstatic/images/service_menu/topup_en_new.gif";

		over[1].src="/retailstatic/images/service_menu/accman_en_over.gif";
		over[2].src="/retailstatic/images/service_menu/cc_en_over.gif";
		over[3].src="/retailstatic/images/service_menu/ft_en_over.gif";
		over[4].src="/retailstatic/images/service_menu/bp_en_over.gif";
		over[5].src="/retailstatic/images/service_menu/cheque_en_over.gif";
		over[6].src="/retailstatic/images/service_menu/ew_en_over.gif";
		over[7].src="/retailstatic/images/service_menu/mp_en_over.gif";
	//	over[8].src="/retailstatic/images/service_menu/ln_en_over.gif";
		over[10].src="/retailstatic/images/service_menu/topup_en_new_over.gif";
}

function MakeArray(n) {
	this.length = n+1;
	for (var i = 0; i<=n; i++) {
		this[i] = new Image();
	}
	return this;
}
if(navigator.userAgent.substring(0,9)>="Mozilla/3") initimages();

function msover(num) {

if(navigator.userAgent.substring(0,9)>="Mozilla/3") {
	if(num==1) AA.src = over[1].src;
	else if(num==2) BB.src = over[2].src;
	else if(num==3) CC.src = over[3].src;
	else if(num==4) DD.src = over[4].src;
	else if(num==5) EE.src = over[5].src;
	else if(num==6) FF.src = over[6].src;
	else if(num==7) GG.src = over[7].src;
//	else if(num==8) HH.src = over[8].src;
	else if(num==10) II.src = over[10].src;
}
}
   
function clear() {
	AA.src = standard[1].src;
	BB.src = standard[2].src;
	CC.src = standard[3].src;
	DD.src = standard[4].src;
	EE.src = standard[5].src;
	FF.src = standard[6].src;
	GG.src = standard[7].src;
//	HH.src = standard[8].src;
	II.src = standard[10].src;
}

function msout(num) {
if(navigator.userAgent.substring(0,9)>="Mozilla/3") {
	if((num==1) && (storedValue !=1)) AA.src = standard[1].src;
	else if((num==2) && (storedValue !=2)) BB.src = standard[2].src;
	else if((num==3) && (storedValue !=3)) CC.src = standard[3].src;
	else if((num==4) && (storedValue !=4)) DD.src = standard[4].src;
	else if((num==5) && (storedValue !=5)) EE.src = standard[5].src;
	else if((num==6) && (storedValue !=6)) FF.src = standard[6].src;
	else if((num==7) && (storedValue !=7)) GG.src = standard[7].src;
//	else if((num==8) && (storedValue !=8)) HH.src = standard[8].src;
	else if((num==10) && (storedValue !=10)) II.src = standard[10].src;
}
}

function clikker(num) {
	clear();
	if (storedValue == 1) AAp.style.display = 'none';
	else if(storedValue == 2) BBp.style.display = 'none';
	else if(storedValue == 3) CCp.style.display = 'none';
	else if(storedValue == 4) DDp.style.display = 'none';
	else if(storedValue == 5) EEp.style.display = 'none';
	else if(storedValue == 6) FFp.style.display = 'none';
	else if(storedValue == 7) GGp.style.display = 'none';
//	else if(storedValue == 8) HHp.style.display = 'none';
	else if(storedValue == 10) IIp.style.display = 'none';	
	if(num == 1) {
		if (storedValue == 1) {
			storedValue = 0;
			setCookie("outline", 0);
		}
		else {
			storedValue = num;
			AAp.style.display='';
			setCookie("outline", num);
			AA.src = over[1].src;
		}
	} 
	else 	if(num == 2) {
		if (storedValue == 2) {
			storedValue = 0;
			setCookie("outline", 0);
		}
		else {
			storedValue = num;
			BBp.style.display='';
			setCookie("outline", num);
			BB.src = over[2].src;
		}
	}
	else 	if(num == 3) {
		if (storedValue == 3) {
			storedValue = 0;
			setCookie("outline", 0);
		}
		else {
			storedValue = num;
			CCp.style.display='';
			setCookie("outline", num);
			CC.src = over[3].src;
		}
	}
	else 	if(num == 4) {
		if (storedValue == 4) {
			storedValue = 0;
			setCookie("outline", 0);
		}
		else {
			storedValue = num;
			DDp.style.display='';
			setCookie("outline", num);
			DD.src = over[4].src;
		}
	}
	else 	if(num == 5) {
		if (storedValue == 5) {
			storedValue = 0;
			setCookie("outline", 0);
		}
		else {
			storedValue = num;
			EEp.style.display='';
			setCookie("outline", num);
			EE.src = over[5].src;
		}
	}
	else 	if(num == 6) {
		if (storedValue == 6) {
			storedValue = 0;
			setCookie("outline", 0);
		}
		else {
			storedValue = num;
			FFp.style.display='';
			setCookie("outline", num);
			FF.src = over[6].src;
		}
	}
	else 	if(num == 7) {
		if (storedValue == 7) {
			storedValue = 0;
			setCookie("outline", 0);
			GGtd.className = '';
		}
		else {
			storedValue = num;
			GGp.style.display='';
			setCookie("outline", num);
			GG.src = over[7].src;
			GGtd.className = 'MSub';
		}
	}

/*	else 	if(num == 8) {
		if (storedValue == 8) {
			storedValue = 0;
			setCookie("outline", 0);
		}
		else {
			storedValue = num;
			HHp.style.display='';
			setCookie("outline", num);
			HH.src = over[8].src;
		}
	}*/

	else 	if(num == 10) {
		if (storedValue == 10) {
			storedValue = 0;
			setCookie("outline", 0);
		}
		else {
			storedValue = num;
			IIp.style.display='';
			setCookie("outline", num);
			II.src = over[10].src;
		}
	}
}
function initclik() {
	if (storedValue == 1) { AAp.style.display = '';AA.src = over[1].src;}
	else if(storedValue == 2) {BBp.style.display = '';BB.src = over[2].src;}
	else if(storedValue == 3) {CCp.style.display = '';CC.src = over[3].src;}
	else if(storedValue == 4) {DDp.style.display = '';DD.src = over[4].src;}
	else if(storedValue == 5) {EEp.style.display = '';EE.src = over[5].src;}
	else if(storedValue == 6) {FFp.style.display = '';FF.src = over[6].src;}
	else if(storedValue == 7) {GGp.style.display = '';GG.src = over[7].src;GGtd.className = 'MSub';}
//	else if(storedValue == 8) {HHp.style.display = '';HH.src = over[8].src;}
	else if(storedValue == 10) {IIp.style.display = '';II.src = over[10].src;}
}
function autoFocus() {
  var x = 0;
  var Fok = false;
		if (document.forms[1]) {
  while ((x < document.forms[1].elements.length) && (!Fok))
   {
     Fok = true;
     if ((document.forms[1].elements[x].type == 'action' ) ||
	(document.forms[1].elements[x].type == 'hidden' ) ||
	(document.forms[1].elements[x].type == 'select-one' ) ||
	(document.forms[1].elements[x].disabled))
     { 
        x++;
        Fok = false;
     }
   }
   if (Fok) { document.forms[1].elements[x].focus(); }
								}
}

function HideShowDetail(){
        if (document.getElementById("dataDetail").style.display == '') 
        {
            document.getElementById("dataDetail").style.display = 'none';
            document.getElementById("SHP").src = "/retailstatic/images/collapse_icon.gif";
            document.getElementById("SHPT").innerHTML = '<a href="javascript:HideShowDetail();">Show details</a>';
        }
        else 
        {
            document.getElementById("dataDetail").style.display = '';
            document.getElementById("SHP").src = "/retailstatic/images/expand_icon.gif";
            document.getElementById("SHPT").innerHTML = '<a href="javascript:HideShowDetail();">Hide details</a>';
        }
}

</script>
</head>
<body bgcolor="#ffffff" text="#000000" link="#333366" alink="#333366" vlink="#333366" onload="init(); init2(); init3();init4(); autoFocus(); " >
<!-- Header starts -->

<table border="0" cellspacing="0" cellpadding="7" width="939" height="20" align="center">
<tbody style="box-shadow: 0 1px 1px rgba(0,0,0,0.19), 0 2px 2px rgba(0,0,0,0.23);">
      <tr bgcolor="#eaeaea">
        <td>
          <div class="top-menu">
	  		<a href="/retail/RetailWelcome.do">
			<img src="/retailstatic/images/bulletm3.png"width="11" height="11" border="0" style="margin: 0px 5px";>

			Main Page</a>
          </div>
        </td>
        <td align="center"  class ='header'>
	   
	            
	            <font color="#606060"><b>
	            N.O.O.N PROPERTY MANAGEMENT 

	            &nbsp;&nbsp;&nbsp;Last login: </b></font>
	    		<font color="#fe0707"><b>24 Nov 2020,8:47 p.m.</b></font>
	            
            
        </td>
        <td align="right">
          <div class="top-menu">
          	
            	<!--<a href="#" onClick="JavaScript:openWebChatWindow()"> Contact Us &nbsp;</a>-->
            	<a href="#" onclick="window.open('http://www.kasikornbank.com/en/ContactUs/Pages/Personal.aspx','','resizable=yes,scrollbars=yes,toolbar=no,location=no,status=yes,menubar=no,width=780,height=450')"> Contact Us &nbsp;</a>
            
          </div>
        </td>
      </tr>
</tbody>
</table>


<table border="0" cellspacing="0" cellpadding="0" width="775" align="center">
   <tr>
	<td><img src="/retailstatic/images/space.gif" width="205" height="6"><kbank:webPath/></td>
  </tr>
</table>

<!--Header ends -->
<table width="939" border="0" cellspacing="0" cellpadding="0" align="center">
  <tr> 
	<td colspan="2"><img src="/retailstatic/images/space.gif" width="16" height="5"></td>
  </tr>
  <tr> 
    <!--td valign="top" width=186-->
    <td valign="top" width=176>



<table width="176" border="0" cellspacing="0" cellpadding="0"  bordercolor="#0000ff">
<tr>
<td align="right" class="left_column"> 

<table width='174' bgcolor='' cellspacing='0' cellpadding='1' border="0" >  	  
  <tr>  
	<td>   
<table border=0 cellpadding=0 cellSpacing=0 width="222">
<tr>  
 <td><img border="0" id="s_menu" src="/retailstatic/images/service_menu/service_menu.gif" width="222"></td>
</tr>
<tr><td>
            <a href="javascript:clikker(1);" >
              <img border="0" id="AA" src="/retailstatic/images/service_menu/accman_en.gif" width="222"></a></td></tr>
			<tr><td class="MSub"><div class="NV2" id="AAp" style="DISPLAY: none">
				<table border="0" cellpadding="0" cellspacing="0" width="222">
					<tr><td class ="MSubAlign"><a class="roll" href="/retail/cashmanagement/inquiry/AccountSummary.do?action=list_domain1">Financial Summary</a></td></tr>
					<tr><td class ="MSubAlign"><a class="roll" href="/retail/cashmanagement/inquiry/AccountSummary.do?action=list_domain2">Account Summary</a></td></tr>
					<tr><td class ="MSubAlign"><a class="roll" href="/retail/accountinfo/AccountStatementInquiry.do">Statement Inquiry</a></td></tr>
					<tr><td class ="MSubAlign"><a class="roll" href="/retail/cashmanagement/AccountDetail.do">Balance Inquiry</a></td></tr>
			        <tr><td class ="MSubAlign"><a class="roll" href="/retail/cashmanagement/TodayAccountStatementInquiry.do">Recent Transactions </a></td></tr>
			        <tr><td class ="MSubAlign"><a class="roll" href="/retail/accountinfo/TransactionList.do?return=true">Online Transactions</a></td></tr>
			        <tr><td class ="MSubAlign"><a class="roll" href="/retail/transaction/recurringtransaction.do">Scheduled Transactions</a></td></tr>
				</table><br>
			</div></td></tr>
<tr><td>
	          <a href="javascript:clikker(2);" >
	            <img border="0" id="BB" src="/retailstatic/images/service_menu/cc_en.gif" width="222"></a></td></tr>
	          <tr><td class="MSub"><div class="NV2" id="BBp" style="DISPLAY: none">
	  			  <table border="0" cellpadding="0" cellspacing="0" width="222">
	  			  	<tr><td class ="MSubAlign"><a class="roll" href="/retail/cashmanagement/inquiry/AccountSummary.do?action=list_domain3&type=c">Card Balance Summary</a></td></tr>
	                <tr><td class ="MSubAlign"><a class="roll" href="/retail/creditcard/CreditCardDetail.do">KBank Card Inquiry</a></td></tr>
	                <tr><td class ="MSubAlign"><a class="roll" href="/retail/creditcard/ccUnbilledTransactionInquiry.do">Unbilled Transaction</a></td></tr>
	                <tr><td class ="MSubAlign"><a class="roll" href="/retail/creditcard/CreditCardTransHistory.do">Card Statement Inquiry</a></td></tr>
	                <tr><td class ="MSubAlign"><a class="roll" href="/retail/creditcard/CreditCardPayment.do">Own KBank Card Payment</a></td></tr>
	                <tr><td class ="MSubAlign"><a class="roll" href="/retail/creditcard/3rdPartyCreditCardPayment.do">Other KBank Card Payment</a></td></tr>
	                <tr><td class ="MSubAlign"><a class="roll" href="/retail/profile/AccountLinkage.do">Add Your Own KBank Card </a></td></tr>

				  </table><br>
	          </div></td></tr>

              </div></td></tr> 
<tr><td>
              <a href="javascript:clikker(3);" >
                <img border="0" id="CC" src="/retailstatic/images/service_menu/ft_en.gif" width="222"></a></td></tr>
              <tr><td class="MSub"><div class="NV2" id="CCp" style="DISPLAY: none">
	            <table border="0" cellpadding="0" cellspacing="0" width="222">
		            <tr><td class ="MSubAlign"><a class="roll" href="/retail/cashmanagement/FundTransfer.do?action=main&accttype=1">Own Account Funds Transfer</a></td></tr>
			        <tr><td class ="MSubAlign"><a class="roll" href="/retail/cashmanagement/FundTransfer.do?action=main&accttype=2">Other Account Funds Transfer</a></td></tr>
		        	<tr><td class ="MSubAlign"><a class="roll" href="/retail/cashmanagement/FundTransfer.do?action=method&accttype=3&itemreq=1&mode=3&sel_base=1&reset=1">Inter-bank Funds Transfer</a></td></tr> 
		        	
		        	<tr><td class ="MSubAlign"><a class="roll" href="/retail/cashmanagement/PromptPayTransfer.do">PromptPay Funds Transfer</a></td></tr>
		     	</table><br>
            </div></td></tr>
<tr><td>
              <a href="javascript:clikker(4);" >
                <img border="0" id="DD" src="/retailstatic/images/service_menu/bp_en.gif" width="222"></a></td></tr>
              <tr><td class="MSub"><div class="NV2" id="DDp" style="DISPLAY: none">
            	<table border="0" cellpadding="0" cellspacing="0" width="222">
	                <tr><td class ="MSubAlign"><a class="roll" href="/retail/cashmanagement/BillPayment.do">Bill Payment</a></td></tr>
					
	                <tr><td class ="MSubAlign"><a class="roll" href="/retail/creditcard/CreditCardPayment.do">Own KBank Credit Card Payment</a></td></tr>
	                <tr><td class ="MSubAlign"><a class="roll" href="/retail/creditcard/3rdPartyCreditCardPayment.do">Other KBank Credit Card Payment</a></td></tr>
				</table><br>
	    	  </div></td></tr>
<tr><td>
  			  <a href="javascript:clikker(10);" >
        		<img border="0" id="II" src="/retailstatic/images/service_menu/topup_en_new.gif" width="222"></a></td></tr>
              <tr><td class="MSub"><div class="NV2" id="IIp" style="DISPLAY: none">
            	<table border="0" cellpadding="0" cellspacing="0" width="222">            	
					<tr><td class ="MSubAlign"><a class="roll" href="/retail/cashmanagement/BillPayment.do?topup=1">Mobile Topup</a></td></tr>
            		<tr><td class ="MSubAlign"><a class="roll" href="/retail/cashmanagement/BillPayment.do?topup=2">Other Topup</a></td></tr>
            		<tr><td class ="MSubAlign"><a class="roll" href="/retail/creditcard/3rdPartyCreditCardPayment.do">Fleet Card Topup</a></td></tr>
				</table><br>
			  </div></td></tr>
<tr><td>
              <a href="javascript:clikker(5);" >
                <img border="0" id="EE" src="/retailstatic/images/service_menu/cheque_en.gif" width="222"></a></td></tr>
        	  <tr><td class="MSub"><div class="NV2" id="EEp" style="DISPLAY: none">
            	<table border="0" cellpadding="0" cellspacing="0" width="222">
	                <tr><td class ="MSubAlign"><a class="roll" href="/retail/cheque/ChequeBookRequest.do">Order Cheque Book</a></td></tr>
	                <tr><td class ="MSubAlign"><a class="roll" href="/retail/cheque/StopCheque.do">Stop Cheque Request</a></td></tr>
	                <tr><td class ="MSubAlign"><a class="roll" href="/retail/cheque/StopChequeStatusEnquiry.do">Stop Cheque Status Inquiry</a></td></tr>
				</table><br>
    		  </div></td></tr>
      
        <tr style="DISPLAY: none"><td><div id="FF"></div><div id="FFp"></div></td></tr>
      
<tr><td>
              <a href="javascript:clikker(7);" >
                <img border="0" id="GG" src="/retailstatic/images/service_menu/mp_en.gif" width="222"></a></td></tr>
              <tr><td id="GGtd" class=""><div class="NV2" id="GGp" style="DISPLAY: none">
              	<table border="0" cellpadding="0" cellspacing="0" width="222">
	              <tr><td class ="MSubAlign"><a class="roll" href="/retail/profile/AccountLinkage.do">Own Account List</a></td></tr>
	              <tr><td class ="MSubAlign"><a class="roll" href="/retail/profile/BeneficiaryList.do">Other's Account List</a></td></tr>
	              
	              <tr><td class ="MSubAlign"><a class="roll" href="/retail/profile/AccountDescription.do">Account Nickname</a></td></tr>
	              <tr><td class ="MSubAlign"><a class="roll" href="/retail/profile/ChangeContactInfo.do">Contact Information</a></td></tr>
	              <tr><td class ="MSubAlign"><a class="roll" href="/retail/profile/MainPageSetting.do">Main Page Setting </a></td></tr>
              	</table><br>
	    	 </div></td></tr>
</table>
<SCRIPT language=JavaScript>initclik();</SCRIPT>

	</td>  
  </tr>  
</table>
 <img src="/retailstatic/images/spacer.gif" height="5" width="5" border="0" /><br>

<table width="176" border="0" cellpadding="0" cellspacing="0">
<tr> 
	<td><img src="/retailstatic/images/service_menu/service_info.gif" border="0"/><br/></td>
</tr>
<tr>
	<td><a href="#service_feature" onmouseover="hideAllmenu();" onclick="window.open('http://www.kasikornbank.com/EN/personal/digital-banking/Pages/k-cyber.aspx','','resizable=yes,scrollbars=yes,toolbar=no,location=no,status=no,menubar=no,width=780,height=450')"><img src="/retailstatic/images/service_menu/service_feature.gif"  border="0" /></a></td>
</tr>
<tr>
	<td><a href="#fees" onmouseover="hideAllmenu();" onclick="window.open('https://www.kasikornbank.com/en/personal/services/payment/Pages/transfer-k-cyber.aspx','','resizable=yes,scrollbars=yes,toolbar=no,location=no,status=no,menubar=no,width=780,height=450')"><img src="/retailstatic/images/service_menu/fees.gif"  border="0" /></a></td>
</tr>
<tr>
	<td><a href="#download_form" onmouseover="hideAllmenu();" onclick="window.open('https://www.kasikornbank.com/en/download/Pages/result.aspx?topic=8&type=27&group=1','','resizable=yes,scrollbars=yes,toolbar=no,location=no,status=no,menubar=no,width=780,height=450')"><img src="/retailstatic/images/service_menu/download_form.gif"  border="0" /></a></td>
</tr>
<tr>
	<td><a href="#faq" onmouseover="hideAllmenu();" onclick="window.open('http://www.kasikornbank.com/EN/personal/digital-banking/Pages/k-cyber.aspx','','resizable=yes,scrollbars=yes,toolbar=no,location=no,status=no,menubar=no,width=780,height=450')"><img src="/retailstatic/images/service_menu/faq.gif"  border="0" /></a></td>
</tr>
<tr>
	<td><a href="#disclaimer" onmouseover="hideAllmenu();" onclick="window.open('https://www.kasikornbank.com/EN/disclaimer/Pages/disclaimer.aspx','','resizable=yes,scrollbars=yes,toolbar=no,location=no,status=no,menubar=no,width=780,height=450')"><img src="/retailstatic/images/service_menu/disclaimer.gif"  border="0" /></a></td>
</tr>
<tr>
	<td><a href="#security_tips" onmouseover="hideAllmenu();" onclick="window.open('http://www.kasikornbank.com/en/security-tips/Pages/index.aspx','','resizable=yes,scrollbars=yes,toolbar=no,location=no,status=no,menubar=no,width=780,height=450')"><img src="/retailstatic/images/service_menu/security_tips.gif"  border="0" /></a></td>
</tr>
</table>

<div id='xbmenu40' class='leftmenu'>  
  <table border='0' cellpadding='0' cellspacing='1' width='158' bgcolor='#808080'>  
	<tr>
	  <td height='20'>
		<a class='leftmenu' href="#" onClick="window.open('/retailstatic/security/CB_Form_ChangeProfile_EN.pdf','','resizable=yes,scrollbars=yes,toolbar=no,location=no,status=no,menubar=no,width=780,height=450')">Changing Profile Request Form</a>
	  </td>
	</tr>
	<tr>
	  <td height='20'>
		<a class='leftmenu' href="#" onClick="window.open('/retailstatic/security/CB_Form_JointAccount_EN.pdf','','resizable=yes,scrollbars=yes,toolbar=no,location=no,status=no,menubar=no,width=780,height=450')">Joint Account Holder Consent Form</a>
	  </td>
	</tr>
	<tr>
	  <td height='20'>
		<a class='leftmenu' href="#" onClick="window.open('/retailstatic/security/CB_Form_Reimbursement_EN.pdf','','resizable=yes,scrollbars=yes,toolbar=no,location=no,status=no,menubar=no,width=780,height=450')">Reimbursement Form to Refund an Incorrect Bill Payment</a>
	  </td>
	</tr>
	<tr>
	  <td height='20'>
		<a class='leftmenu' href="#" 
			onClick="window.open('/retailstatic/security/Claim_Letter_for_Web_Card.pdf','','resizable=yes,scrollbars=yes,toolbar=no,location=no,status=no,menubar=no,width=780,height=450')">Claim Letter for K-Web Shopping Card</a>
	  </td>
	</tr>
	<tr>
	  <td height='20'>
		<a class='leftmenu' href="#" 
			onClick="window.open('/retailstatic/security/KCyberBanking_AddOtherAcctForm_EN.pdf','','resizable=yes,scrollbars=yes,toolbar=no,location=no,status=no,menubar=no,width=780,height=450')">Add Other Account Form</a>
	  </td>
	</tr>
	<tr>
	  <td height='20'>
		<a class='leftmenu' href="#" 
			onClick="window.open('/retailstatic/security/KCyber_AddOwnAccForm_EN.pdf','','resizable=yes,scrollbars=yes,toolbar=no,location=no,status=no,menubar=no,width=780,height=450')">Add Own Account Form</a>
	  </td>
	</tr>

  </table> 
</div>  
<div id='xbmenu41' class='leftmenu'>  
  <table border='0' cellpadding='0' cellspacing='1' width='158' bgcolor='#808080'>  
	<tr>
	  <td height='20' bgcolor='#dfdfdf'>
		Will be available soon
	  </td>
	</tr> 
  </table> 
</div> 

</td>
</tr>
</table>

  </td>

<td valign="top" width="100%">
<table width="100%" border="0" cellspacing="0" cellpadding="0" align="center">
	   







		
<tr>

		   <tr>
			 <td><img src="/retailstatic/images/spacer.gif" height="10" width="5" border="0" /></td>
		   </tr>






















<SCRIPT language="JavaScript" type="text/javascript">
<!--

function OpenDetail(FileName,WindowName){
	WindowFeature="Resizable=Yes,ScrollBars=Yes,MenuBar=No,Directories=No,ToolBar=No,Location=No,Status=No,Width=800,Height=500,ScreenX=0,ScreenY=0,Top=0,Left=0"
	
	newWindow=open(FileName,WindowName,WindowFeature);
	if (newWindow.opener == null) { newWindow.opener = self; }
		if (newWindow.focus) { newWindow.focus(); }
}

function gotoURL(targ){ //v3.0
  if (targ != null)
	  eval("window.location='"+ targ + "'");
}

//-->
</SCRIPT>

<style type="text/css">
A.announcement:link {
	COLOR: #269b11;
	FONT-WEIGHT: bold;
	FONT-FAMILY: verdana;
	FONT-SIZE: 11px
}

A.announcement:visited {
	COLOR: #269b11;
	FONT-WEIGHT: bold;
	FONT-FAMILY: verdana;
	FONT-SIZE: 11px
}

A.announcement:hover {
	COLOR: #269b11;
	FONT-WEIGHT: bold;
	FONT-FAMILY: verdana;
	FONT-SIZE: 11px
}

A.announcement:active {
	COLOR: #269b11;
	FONT-WEIGHT: bold;
	FONT-FAMILY: verdana;
	FONT-SIZE: 11px
}

A.announcement:unknown {
	COLOR: #269b11;
	FONT-WEIGHT: bold;
	FONT-FAMILY: verdana;
	FONT-SIZE: 11px
}

li {
	padding: 4px;
}
</style>

<form name="RetailWelcomeForm" method="post" action="/retail/RetailWelcome.do"><div><input type="hidden" name="org.apache.struts.taglib.html.TOKEN" value="0.21369009104596925"></div>
	<table width="100%" border="0" cellspacing="0" cellpadding="0"
		onmouseover="hideAllmenu();">
		<tr>
			<td width="10"><img src="/retailstatic/images/space.gif"
				width="10" height="16" alt=""></td>
			<td>
			<table width="100%" border="0" cellspacing="0" cellpadding="5">

				<!------------ Notice ------------------>
				
				<!------------ Announcement ------------------
				
				<tr>
					<td colspan="2">
					<h1>Announcement(s)/ Information</h1>
					</td>
				</tr>
				<tr>
					<td style="padding-left:30px" colspan="2" width="580px">
					
                  	 	<li><a href="#" class="announcement" onClick="javascript:OpenDetail('','Announcement')"></a>&nbsp;<img src="/retailstatic/images/new.gif"/></li>
			
					</td>
				</tr>
				<tr>
					<td colspan="2"><img src="/retailstatic/images/pixel_gray.gif"
						width="100%" height="2" alt=""></td>
				</tr>
    -->
				<!------------ Alert & Promotion ------------------>
				<!-- 		
				<tr>
					<td colspan="2">
					<h2>My recent Online Transaction(s)</h2>
					</td>
				</tr>
				<tr>
					<td style="padding-left:30px" colspan="2">
					<li><a
						href="javascript:OpenDetail('','AlertTransaction')">Funds
					Transfer/Bill Payment -> Transaction List </a></li>
					</td>
				</tr>
				<tr>
					<td colspan="2"><img
						src="/retailstatic/images/pixel_gray.gif"
						width="100%" height="2" alt=""></td>
				</tr>
				        -->

				<!----------- Favorite transaction ----------->
				

				<!------------------------ Currency Rates ---------------------------->
				
				<!-------------------------- Interest Rates -------------------------->
				
				<tr>
					<td colspan="2">
					<h1>Interest Rates for an Individual</h1>
					</td>
				</tr>
				
				<tr>
					<td colspan="2"></td>
				</tr>
				<tr>
					<td>
					<h2>Savings</h2>
					</td>
					<td align="right"><b>Last Updated : </b>14 Jul 2020</td>
				</tr>
				<tr>
					<td colspan="2">
					<table width="100%" border="1" cellspacing="0" cellpadding="0"
						bordercolor="#FFFFFF">
						<tr>
							<td class="table2_left" valign="bottom" width="50%">Type</td>
							<td class="table2_center" width="50%">Rate</td>
						</tr>
						
						<tr>
							<td width="50%" class="inner_table_left">< 1,000 million</td>
							<td width="50%" class="inner_table_center">0.25</td>
						</tr>
						
					</table>
					</td>
				</tr>
				
				
				<tr>
					<td colspan="2"></td>
				</tr>
				<tr>
					<td>
					<h2>Fixed Deposit</h2>
					</td>
					<td align="right"><b>Last Updated : </b>14 Jul 2020</td>

				</tr>
				<tr>
					<td colspan="2">
					<table width="100%" border="1" cellspacing="0" cellpadding="0"
						bordercolor="#FFFFFF">
						<tr>
							<td rowspan="2" class="table2_left" valign="bottom" width="33%">Type</td>
							<td class="table2_center" colspan="4">Rate</td>
						</tr>
						<tr>
							<td width="17%" class="table3_center">3 Months</td>
							<td width="17%" class="table3_center">6 Months</td>
							<td width="17%" class="table3_center">12 Months</td>
							<td width="17%" class="table3_center">24 Months</td>
						</tr>
						
						<tr>
							<td width="33%" class="inner_table_left">< 3.0 million</td>
							<td width="17%" class="inner_table_center">0.37</td>
							<td width="17%" class="inner_table_center">0.45</td>
							<td width="17%" class="inner_table_center">0.45</td>
							<td width="17%" class="inner_table_center">0.5</td>
						</tr>
						
						<tr>
							<td width="33%" class="inner_table_left">3.0 million to < 10.0 million</td>
							<td width="17%" class="inner_table_center">0.37</td>
							<td width="17%" class="inner_table_center">0.45</td>
							<td width="17%" class="inner_table_center">0.45</td>
							<td width="17%" class="inner_table_center">0.5</td>
						</tr>
						
						<tr>
							<td width="33%" class="inner_table_left">10.0 million to < 30.0 million</td>
							<td width="17%" class="inner_table_center">0.37</td>
							<td width="17%" class="inner_table_center">0.45</td>
							<td width="17%" class="inner_table_center">0.45</td>
							<td width="17%" class="inner_table_center">0.5</td>
						</tr>
						
						<tr>
							<td width="33%" class="inner_table_left">30.0 million to < 50.0 million</td>
							<td width="17%" class="inner_table_center">0.37</td>
							<td width="17%" class="inner_table_center">0.45</td>
							<td width="17%" class="inner_table_center">0.45</td>
							<td width="17%" class="inner_table_center">0.5</td>
						</tr>
						
						<tr>
							<td width="33%" class="inner_table_left">50.0 million to < 100.0 million</td>
							<td width="17%" class="inner_table_center">0.37</td>
							<td width="17%" class="inner_table_center">0.37</td>
							<td width="17%" class="inner_table_center">0.45</td>
							<td width="17%" class="inner_table_center">0.5</td>
						</tr>
						
						<tr>
							<td width="33%" class="inner_table_left">100.0 million to < 5000.0 million</td>
							<td width="17%" class="inner_table_center">0.37</td>
							<td width="17%" class="inner_table_center">0.45</td>
							<td width="17%" class="inner_table_center">0.45</td>
							<td width="17%" class="inner_table_center">0.5</td>
						</tr>
						
						<tr>
							<td width="33%" class="inner_table_left">> 5000.0 million</td>
							<td width="17%" class="inner_table_center">0.37</td>
							<td width="17%" class="inner_table_center">0.45</td>
							<td width="17%" class="inner_table_center">0.45</td>
							<td width="17%" class="inner_table_center">0.5</td>
						</tr>
						
					</table>
					</td>
				</tr>
				
				
				<tr>
					<td colspan="2">
					<div align="center"><a
						href="/retail/profile/MainPageSetting.do"><img
						src="/retailstatic/images/button/Personalize_en.gif" border="0"></a>
					</div>
					</td>
				</tr>
				<tr>
					<td colspan="2"><img src="/retailstatic/images/pixel_gray.gif"
						width="100%" height="2" alt=""></td>
				</tr>
				
				<!-------------- Foreign Currency Deposit Rates ---------------------->
				
			</table>
			</td>
		</tr>
	</table>
</form>
	  <td align="center" valign="top" width="115">
		<table width="94" border="0" cellspacing="0" cellpadding="0">
			<tr>
				<td>
		        <a href='/retail/accountinfo/AccountStatementInquiry.do'><img src="/retailstatic/images/right_menu/StatementInquiry_en.gif" vspace="2" border="0"></a><br>
                <a href='/retail/cashmanagement/AccountDetail.do'><img src="/retailstatic/images/right_menu/BalanceInquiry_en.gif" vspace="2" border="0"></a><br>
				<a href='/retail/cashmanagement/FundTransfer.do?action=main&accttype=1'><img src="/retailstatic/images/right_menu/OwnAccountFundsTransfer_en.gif" vspace="2" border="0"></a><br>
				<a href='/retail/cashmanagement/FundTransfer.do?action=main&accttype=2'><img src="/retailstatic/images/right_menu/OtherAccountFundsTransfer_en.gif" vspace="2" border="0"></a><br>
				<a href='/retail/cashmanagement/BillPayment.do'><img src="/retailstatic/images/right_menu/BillPayment_en.gif" vspace="2" border="0"></a><br>
			    </td>
			</tr>











<!--Content end here-->
	  </table>
	  </td>
	</tr>
		
<!--Body and right columm end-->
     </td>
  	</tr>
  </table>
 </td>
 </tr>
 
 <tr> 
   <td>
	  <img src="/retailstatic/images/space.gif" width="16" height="16">
   </td>
 </tr>
</table>
<!-- Footer starts here -->
<div class="footer">
<table border="0" cellspacing="0" cellpadding="0" align="center" width="939" height="17" align="center" onmouseover="hideAllmenu();" >
  <tr>
    <td align="center" class="footer">
	  <div align=center><img src="/retailstatic/images/kBank_colorbar.jpg" width="100%"></div>
	  <div class="list1"><script type="text/javascript">document.write(COPY_RIGHT);</script></div>
    </td>
  </tr>
</table>
</div>
<!--Footer ends here -->

</body>
</html>

`
