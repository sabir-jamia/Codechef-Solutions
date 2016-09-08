<?php 
fscanf(STDIN, "%d\n", $t);
for($i = 0; $i < $t; $i++) {
		fscanf(STDIN, "%s\n", $str); 
		$flag = true;
		$count = strlen($str);

		if($count % 2 != 0 ) {
			$len = ($count +1)/ 2;
		} else {
			$len = $count / 2;
		}

		for ($j=0; $j < $len; $j++) { 		
			if($str[$j] == $str[$count - $j - 1]) {
				if($str[$j] == '.') {
					$str[$j] = $str[$count - $j - 1] = 'a';
				}
			} elseif($str[$j] == '.') {
				$str[$j] = $str[$count - $j - 1];
			} elseif ($str[$count - $j - 1 ] == '.') {
				$str[$count - $j - 1] = $str[$j];
			} else {
				$flag = false;
			}
		}

		if ($flag) { 
			printf("%s\n", $str); 
		} else { 
			printf("%s\n", "-1");
		} 
}