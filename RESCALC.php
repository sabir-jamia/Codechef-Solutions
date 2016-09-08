<?php 
main();
function main() {
	fscanf(STDIN, "%d\n", $t);
	for($i = 0; $i < $t; $i++) {
		fscanf(STDIN, "%d\n", $numPlayers); 
		for ($j=0; $j < $numPlayers; $j++) { 
			fscanf(STDIN, "%d\n", $numCookies);
			$points[$j] = $numCookies;
			$cookies = array();
			print_r($numCookies);
			for ($k=0; $k < $numCookies; $k++) { 
			 	fscanf(STDIN, "%d\n", $cookie);
			 	print_r($cookie);die; 
			 }
			 print_r($cookies);die;
			$points[$j] += count($cookies);
		}
		$maxIndex = maximum($points);
		if($maxIndex == 0) {
			printf("%s\n", "chef");	
		} else {
			printf("%d\n", $maxIndex + 1);
		} 
	}
}

function maximum($points) 
{
	$len = count($points);
	$index = 0;
	$max = $points[$index];
	for($i = 1; $i < $len; $i++) {
		if($max < $points[$i]) {
			$index = $i;
			$max = $points[$index];
		}
	}
	return $index;
}