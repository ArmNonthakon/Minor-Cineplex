
export const ToUpperFirstLetter = (input:string)=>{
    let text = input
    let fText = text.charAt(0).toUpperCase()
    const newText = fText + text.slice(1)
    return newText
}