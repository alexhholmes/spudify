// Cookie reading function from https://www.w3schools.com/js/js_cookies.asp
const COOKIE_USER_ID = "xzczxc302"
const COOKIE_USER_AUTH = "malscmxc2"

function setCookieWithMonthExpiry(name, value) {
    let expiry = new Date();
    expiry.setMonth(expiry.getMonth() + 1); // Adds one month to current date
    document.cookie = `${name}=${value};expires=${expiry.toUTCString()}`;
}

// Deletes cookie by forcing expiration
function deleteCookie(name) {
    let expiry = "Thu, 01 Jan 1970 00:00:01 GMT"
    document.cookie = `${name}='';expires=${expiry}`;
}

function getCookie(cname) {
    let name = cname + "=";
    let decodedCookie = document.cookie;
    let ca = decodedCookie.split(';');
    for(let i = 0; i <ca.length; i++) {
        let c = ca[i];
        while (c.charAt(0) === ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) === 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

function checkCookie(cookieName) {
    let cookieVal = getCookie(cookieName);
    return cookieVal !== "";
}

export {
    COOKIE_USER_ID,
    COOKIE_USER_AUTH,
    setCookieWithMonthExpiry,
    getCookie,
    checkCookie,
    deleteCookie
}