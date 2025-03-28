import { ref } from 'vue'
export const dark_mode = ref(false)
export const current_view = ref('home')
export const current_account = ref('Microsoft')
export const current_download = ref('Auto-Install')
export const current_about = ref('Treasure-Box')
export const current_setting = ref('Game')
export const allowCookie = ref(false)

export function setCookie(key: string, value: any) {
    let date = new Date()
    date.setTime(date.getTime() + (30*24*60*60*1000))
    document.cookie = key + "=" + value + "; expires=" + date.toUTCString() + "; path=/"
}
export function getCookie(key: string): any {
    let cookieArray = document.cookie.split("; ")
    for (let i = 0; i < cookieArray.length; i++) {
        let cookie = cookieArray[i].split("=")
        if (cookie[0] === key) {
            return cookie[1]
        }
    }
    return ""
}