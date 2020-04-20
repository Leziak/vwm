import React from 'react'

export default function (props) {
  const {imageSrc, height, width, id} = props
  return (
    <div onClick={() => console.log('hello')} style={{ margin: '20px' }}>
      <img src={imageSrc} alt='Arryn' height={height} width={width} />{' '}
    </div>
  )
}
