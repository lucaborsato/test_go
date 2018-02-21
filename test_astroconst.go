// test of astrconst
package main

import (
	"fmt"
	acst "github.com/lucaborsato/astroconst"
)

func main() {
	fmt.Printf("TwoPi  = %23.16e\n", acst.TwoPi)
	fmt.Printf("d2r    = %23.16e\n", acst.Deg2Rad)
	fmt.Printf("r2d    = %23.16e\n", acst.Rad2Deg)
	fmt.Printf("180*d2r= %23.16e\n", 180.*acst.Deg2Rad)
	fmt.Printf("Pi*r2d = %23.16e\n", acst.Pi*acst.Rad2Deg)
	
	fmt.Printf("Msun2ear = %23.16e\n", acst.Msun2ear)
	fmt.Printf("Mear2sun = %23.16e\n", acst.Mear2sun)
	fmt.Printf("Msun = %23.16e\n", acst.Msun)
	fmt.Printf("Mear = %23.16e\n", acst.Mear)
	fmt.Printf("Mjup = %23.16e Mear\n", acst.Mjup2ear)
	fmt.Printf("Mnep = %23.16e Mear\n", acst.Mnep2ear)
}
