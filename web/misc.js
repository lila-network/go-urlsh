function Logout() {
    document.cookie = 'gourlsh_auth=; Max-Age=-1; path=/; domain=' + location.hostname;
    document.location = "/admin/login"
}