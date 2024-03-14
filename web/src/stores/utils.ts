
export const parseFecha = (fecha:string, ful:boolean) => {
  const partes = fecha.split('T')
  const fs = partes[0].split('-')
  let date = `${fs[2]}/${fs[1]}/${fs[0]}`
  if(ful){
    const fx = partes[1].split(':')
    date += ` ${fx[0]}:${fx[1]}`
  }
  return date
}
