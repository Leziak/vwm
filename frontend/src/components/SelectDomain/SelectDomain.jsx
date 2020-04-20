import React from 'react'

import Item from './Item/Item'

import Arryn from './../../images/heraldry/Arryn.png'
import Baratheon from './../../images/heraldry/Baratheon.png'
import Greyjoy from './../../images/heraldry/Greyjoy.png'
import Lannister from './../../images/heraldry/Lannister.png'
import Martell from './../../images/heraldry/Martell.png'
import Stark from './../../images/heraldry/Stark.png'
import Targaryen from './../../images/heraldry/Targaryen.png'
import Tully from './../../images/heraldry/Tully.png'
import Tyrell from './../../images/heraldry/Tyrell.png'

const containerStyle = {
  display: 'flex',
  flexDirection: 'row',
  justifyContent: 'center',
  flexWrap: 'wrap',
  width: '800px',
  margin: '0 auto',
}

const itemStyle = {
  margin: '20px',
}

const itemHeight = 180
const itemWidth = 180

export default function (props) {
  return (
    <>
      <h2 style={{textAlign: 'center'}}>SELECT DOMAIN</h2>
      <div style={containerStyle}>
        <Item id="arryn" height={itemHeight} width={itemWidth} imageSrc={Arryn} />
        <Item id="baratheon" height={itemHeight} width={itemWidth} imageSrc={Baratheon} />
        <Item id="greyjoy" height={itemHeight} width={itemWidth} imageSrc={Greyjoy} />
        <Item id="lannister" height={itemHeight} width={itemWidth} imageSrc={Lannister} />
        <Item id="martell" height={itemHeight} width={itemWidth} imageSrc={Martell} />
        <Item id="stark" height={itemHeight} width={itemWidth} imageSrc={Stark} />
        <Item id="targaryen" height={itemHeight} width={itemWidth} imageSrc={Targaryen} />
        <Item id="tully" height={itemHeight} width={itemWidth} imageSrc={Tully} />
        <Item id="tyrell" height={itemHeight} width={itemWidth} imageSrc={Tyrell} />
      </div>
    </>
  )
}
