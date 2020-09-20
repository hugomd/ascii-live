package frames

var Rick = DefaultFrameType(rickFrames)

var rickFrames = []string{
	`tfffLfttffffffLLffffttttttttttt1tfffffffffftttttt111111tffffffffftttff
fttfffttfffffffffttfftttttt1t1tt11tfffffftttttt111t111111ttffffLLffttf
LfttttffLLLfttttftfLftttttttffffttttffffttttfftt11t111tt111ttfffLffttf
LLffttfLLffttfffftfLfttttttffttt11111ttttffffffftt1111tft11111ttfffftt
LLLfftffttttffLffftfftttttft1i,,::::,iitffffffffftt111tft1tttt1ttffLft
LLLfttttffftffffftfttttttff1:,,,,,,,,::itfffffffft11111111ttttttttttt1
LLLfttfftffffffffLLLftttttt,::,:::,,,,::1tfffftfft11111ttt1ttt1tfttt11
LffftfLfffttfLLffLLLfttttti,,,iii11111,:i11ttffft111111tfft11t11tt1ttt
fttttfffffttLLLLfLLLftttfti:,,,iiiiiiii,it111tt111t1111tffttfftt11ttft
ttttttttttttfLLLffLLfttttt1,,,,,,iiiiii,1tft11tt1111111tfftffffttttttt
ffffffffttffttLLfffLfttttt11i,,,,iiiiiii1tt11tfftt11111ttt1fffttt1tttt
ffffffffttffftffttftttttff1ii,,,,iii1iii111tt1ttttt1111t111ffttffttfff
fffffffttttttttttffftttttft1ii,,,ii1iii111tfft11ttt111111t1tt1tfft1ttf
fffffftttttttt11fffftttt1t111iiii,iiiii11ttffft11111111ttttttt1tt11ttf
fttttttttffttt11tffttttttt11ttii,,,iiii11i1ffftt1111111tttt111tttt11tt
tttttttfffftttt11tfftttttt111t1i,,,iiii1i,,:,i1111t1111ttt111t1tfftt11
fffftttffffttffttttft111tt11111i,,,iiii1,,,....,:::,i11tt111tt11tffftt
ffffttttfffttffttttft1111111111,,,,,,,it,,...........,:it111ttt1tffftt
ffft111tffttfttttttt1111t1i,:,11i1i,,1t1:,,............,itt11tt1tfftt1
ttt1ttt11ttttttftt111111i:,..:1111iii11,,...............,tt111t11ttttt
ttttffftt111ttffftt11i,,.,...,iiii,,,,:,,...............:1ttt1111ttfff
11ttttttttttfffffft1:,..........::::::,,,...............,1ttttt11ttfff
11ttttttttttfffffft1:...........:,::,:..................,1tttttt1ttfff
1ttftttfftttttffft11:...........,:::::..................,1tttft11tffff
1tttttttttttttttttt1,...........,::::,...................,t1t1111tffff
1tttttttt1tfftttttt1:...........,::::,...... ............,1111tt11ffff
1tttttttt1tffftttft1:...........,:::,,....................it1ttt11tfff
1ttttttttttffttttttt:...........,::,,.....................:111tt11tttt
1tftttttftttttttttt1,,:,:.......,:,,,....................,,1111t11tttt
1ttffffffttttfffft11ii1i:.......,:,,,.....................,1tt1111tttt
1tttttttttttttttttti,,,,:.......,:,,,....... .............,itt1111tttt
1111111ttttttttttt1i,,,,:.......,:,,,......................itt11111111
1111111ttttttttttt1,,,,,:...... ,:,,,,.....................itt11t11111
1111111tttttttttt1,:,::,,...... ,:...,.....................ittttt11111
111111111tttttttt1i,.....  .... .:........................,i1tttt11111
1111111111t11t11111,,....,,.... ,:::,,,...................,11ttttt11t1`,

	`tfffLLttffffffLLffffttttttttttt1tffffffffffttttttt11111tffffffffftttff
fttfffttfffffffffttfftttttt1tttt1ttftfffftt1ttt111t111111ttffffLLffttf
LfttttffLLLfttttttfLftttttttfft11i,,:,,i11ttfftt11t111tt111tttfffffttt
LLffttfLLLfttfffftfLLttttttfff1,,,,,,,,,:,tffffft11111tft11111ttfffftt
LLLfftffttttffLffftfftttttfft1,:,,,,,::::,,tffftttt111tft1tttt1ttffLft
LLLfttttffftffffftfftttttfffti:,,,iiii11i,:tffffft11111111ttftttttttt1
LLLftffftfLftfftfLLLftttttfft::i,iiiiii11i:1ffftft11111ttt1ttt1ttttt11
LLfftfLLfffffLfffLLLfttttt1tt,:,,,,,iiiiiii1tffft11111tffft11t11tt1ttt
fttttfffffttLLLLfLLLftttft111i,,,,,,iiiiii111tt111t1111tffttfftt11tfft
ttttttttttttfLfLffLLfttttttt1iii,,,,,ii11i1t11t11111111tffttfffttttttt
ffffffffttfftfLLfffLftttttfft1ii,,,,iii1i1tt1tfftt11111ttt1tffttt1ttt1
fffffffttfffftffffftttttfffft111,,,,iiii1111t1ttttt1111tt11ffttfftttff
fffffffttttttttttffftttttttt11ttii,iiiiii11tft11ttt111111t11t1tfft1tff
fffffftttttttt11ffffttttt1111tftiii,,iiii1,,,1111111111tttt11111111ttf
fttttttttffttt11tffftttttt11ttf1,,,i,iiii1:,,.,,::,ii11tttt111tttt11tt
tttttttfffftttt11tfftttttt111tt1,,,i,,ii1t:,,........,:,1t111t1tfftt11
fftftttffffttffttttftt11tt111i,,,,,,,,,it1,,............,,11tt11tfffft
fffft1ttffftfffttttft111111,:,.,iii,,,i11,.,.............,1tttt1tffftt
ffftt11tffttfttttttt1111i,,....,1iii,ii,,................,111tt1tfft11
ttt1tttt1ttttttttt1111,:,.......i1,,,:::.................,it11t11ttttt
111ttfftt11ttfffftt1ti..........:i,::,:,..................ittt111t1tff
11tttttttt1tffffffftti,.........,::::::,..................itttt11ttfff
11ttttttttttffffffft1i,.........,:,:::,...................,1tttt11tfff
11tftttfftttttffft1111:.........,:,:::,....................,tt111tffff
1tttttttftttttttttt111,.........,:,::,,....................,111111ffff
11tttttttttffttttfft11i.........,:.:,,......................:1tt11ffff
11ttttttt1tffftttfft111:....... ,:.:,,.......................1tt11tfff
1ttttttttttfttttttft111i,...... ,:,,,,......................,itt11tttt
1tfttttfftttttffttt1111i,...... ,:.,,,......................:11111tttt
1tfffffffttttfffft11111,....... ,,.,,....... ...............it1111tttt
1ttttttttttttttttt1111,,....... ,,.,,......................:tt11111ttt
11t1111tttttttttt1iiii:,,...... ,,.........................,t111111111
1111111ttttttttt1,,iiii:....... ,,............ ...........,tttt1111111
1111111ttttttttt1,,,iii,... ....,::::,....,:::,,.........,1ttttt111111
111111111tttttttt1i,,ii:  .,....,:::::... .,,,,,:,.....,,1ttttttttt111
1111111111tt1111111i,,.,:,i: ...,:::::.....:,,,,,:.....,1ttttttttt11t1`,

	`tfffLLttffffffLLffffttttttttttt1ttt1iiiii1tttttttt11111tffffffffftttff
fttffftffffffffffttffttttt11tttt1i::,,,,,,:,111111t111111ttffffLLffttf
LfttttffLLLfftttttfLftttttttfft1i,,,,,,,,,,:,ttt11t111tt111tttfffffttt
LLffttfLLLfttfffftfLLtttttffffti,::,,,,,iii::1ftt11111tft11111ttfffftt
LLLLftffttttffLffftffttttffffft,:,iiii11111i:iftttt111ttt1tttt11tffLft
LLLfttttffftffLLftfftttttffffft,:,,,,,iiiiii,1fftt11111111ttftttttttt1
LLLftffftfLftffffLLLftttttffffti,,,,,,,ii,iiitffft11111ttt1ttt1tfttt11
LLffffLLfffffLfffLLLLtttttttff1iii,,i,,iii111tfft111111ffft11t11tt1ttt
fttttfffffttLLLLfLLLffttftt11t1iii,,,,,iiiii1tt111t1111tffttfftt11tfft
ttttttttttttfLLLffLLftttttttft111i,,,,iiii1111111111111tffttfffttttttt
ffffffffttfttfLLfffLftttttfffft11i,,,,iiii1t1ttttt11111ttt1tffttt11ttt
fffffffttfffftffffftttttfffft11t1i,,,,,iii1ii1ttttt1111t111ffttfft1tff
fffffffttttttttttffftttttttt11tft1i,,,,iii1i,:,i1ttt11111t11t1tfft1ttf
fffffftttttttt11ffffttttt1111tfti,ii,,iiii1,,,...,:,ii1tttt1111tt11ttf
fttttttttffttt11tffftttttt11t1iii,,ii,,,i1t:,.........,:,1t111ttt111tt
tttttttfffftttt11tfftt11tt1i:,.:,,,i,,,itti,,............,,1111ttttt11
fffftttffffttffttttft1111i:....,,,,ii,i11,.,,.............:1tt1ttfffft
ftffttttffftfffttttft1i:,.......:,,,,,,,:,................:tttt1tffftt
ffftt11tffttftttttttt1,.........,ii,,,,,,.................:11tt1tfft11
ftt1ttt11ttttttttt111i,.........,11,:,,:,.................,1t1t11ttttt
111ttfftt11ttffffftt11,.........,i,::::,..................,1tt11111fff
11tttttttt1tffffffft11,..........::::::,...................,tt1111tfff
11ttttttttttffffffft11:..........:::::,..,,:,...............ittt1ttttt
11tftttffttttfffft1111:..........:::::,..,1i,:,..............1t11tffff
1tttttttftttttttttt111,..........::,:,...,ii,,,,.............:1111tfff
1ttttttttttfftttttft11i,.........::,,,...,,,,,i:..............,111tfff
11tttttttttffftttfft11i,.........::,,,...,,,,,i:..............:111tfff
1ttttttttttfftttttft111:.........:,,,....::,,,,:.............,1t11tttt
1tfttttfftttttfftttt111:.........:,,,......,::::..,.........,i1111tttt
1tfffffffttttfffft1111i:,........:,,,..... .......,,...... .it1111tttt
1ttttttttttttttttt1iiii,,........:...........  ...........:itt11111ttt
11tt111tttttttttti,,,iii:... ....:,,...................:i1tttt11111111
1111111tttttttttt1,,,,,,,. .,....::::..................:1tttttt1111111
1111111ttttttttttt1i,,::..:,1,...::::...................,ttttttt111111
1111111111tttttttt111i,.,it1,....::::,..................,1ttttttttt111
1111111111tttt11111111ii1t1i:...,::::,........,..........itttttttt11t1`,

	`tfffLLttffffffLLffffttttttttttt1tt1::,,,,,,:,1ttt111111tfftfffffftttff
fttffftffffffffffttfftttttt1ttt11i,:,,,,,,,,,:1111t1111111tffffLLffttf
LfttttffLLLfftttttfLftttttttffft1,:::,,,,,,,::it11t111tt111tttfffffttt
LLffttfLLLfttfffftfLLtttttfffffti:,,ii111111i:itt11111tft11111ttfffftt
LLLLftffftttffLffftffttttfffffft,:,,,,,iiiiii,ittt1111ttt1tttt11tffLft
LLLfttttffftffLLftfftttttffffffti,,,,,,,ii,iiitttt11111111ttftttttttt1
LLLftffftfLftffffLLLftttttfffff1iii,,i,,i1i1i1tttt11111ttt1ttt1tfttt11
LLffffLLfffffLfffLLLftttttttfft1i,,,,,,,iiiiitfft111111ffft11t11tt1ttt
fttttffffftfLLLLfLLLftttftt1111111i,,,,iiii111t111t1111tffttfftt11tfft
ttttttttttttfLLLffLLftttt1tttt11tti,,,iiiii111111111111tffttfffttttttt
ffffffffttfttfLLfffLftttttffft1111i,,,,,ii111tfttt11111ttt1tffttt11ttt
ffffffffttffftffffftttttfffft11t11i,,,,iii1i:,1ttft11111111ffttfftttff
fffffffttttttttttffftttttft111tttiii,,iiii1,,,,:,i1111111111t1tfft1ttf
ffffffttttttttttfffftttt1111i,:,iiii,,,i,11:,,.....,:,i1tft1111tt111tf
fttttttttfftt111tfffttt11i,,,..:i,,,i,ii1t1,,..........,:i1111ttt1111t
tttttttffffttttt1tffti,:,,.....,,,,iiii1ti,,,.............,1111ttttt11
fffftttffffttffttttf1:.........,,:,,,,,,,,.,,,............:1tt1ttffff1
ftfft1ttffftfffttttf1,..........,,,,,,,,:.:ii,,:,.........:tttt1tffftt
ffftt11tffttfttttt1t1,...........ii,,,::,,,i,,ii:.........:tttt1tfft11
ftt1tttt1ttttttttt11i,..........,1i,,::,.:i,,,,i:.........,1t1t11ttttt
11tttffttt1tttffft11i,..........,ii,:::,.,:,,iii:..........it111111tff
11tttttttt1tfffffftti,...........:::::,....:,,ii:..........:tt111ttfff
11ttttttttttfffffft1i,...........:::::,.....,:,,:.... ......,1t111tttt
11tftttffttttfffft111:...........:::::,........,..,..........,1111tfff
1tttttttftttttttttt11:...........::,:,........ ...:,,.........:111tfff
1ttttttttttffttttttti,...........:::,,.............,...........,i1tfff
11tttttttttffftttfft,,:,:........:::,.............. ...........:11tfff
1ttttttttttfftttttftii1i,:.......:::,.......... ..............:t111ttt
1tfttttfftttttffttt1ii,ii,,......::,,........................:1111tttt
1tfffffffttttfffft11i,,ii,,......:,,,..........:........   .,t1111tttt
1tttttttttttttttttt11,:,i:,,.....:,.................. .,i,,1tt1111tttt
11t111ttttttttttttt11i:,,:1,....,::....................:1ttttt11111111
111111tttttttttttttt11i,,1t:....:::::...................,tttttttt11111
1111111ttttttttttttt111ttt1,,,.,:::::,..................,1ttttttt11111
111111111tttttttttt1111111i,::.,,::,:,...................,ttttttttt111
1111111111ttttttt11111111t,.,,.,,,,,,,...................,1ttttttt11t1`,

	`tfffLLttffffffLLffffttttttttttt1ttt,:,,,,,,,:itttt11111tfftfffffftttff
fttffftffffffffffttfftttttt1tttt11i:,,,,,,,,,:i111t1111111tffffLLffttf
LfttttffLLLfftttttfLftttttttffft1,::,,,,,,,,,,,t111111tt111tttfffffttt
LLffttfLLffttfffftfLLtttttffffff1,:iiii111111,:tt11111tft11111ttfffftt
LLLLftffftttffLffftffttttfffffffi:,,,,,iiiii1,,ttt1111ttt1tttt11tffLft
LLLfttttffftffLLftfftttttfffffff1,,,,,,,ii,,ii1ttt11111111ttftttttttt1
LLLffffftfLftffffLLLftttttfffff1iiiiii,,iiii1itttt11111ttt1ttt11tttt11
LLffffLLftffLLLffLLLftttttttffftiii,,,,,ii1ii1tft111111ffft11t11tt1ttt
fttttfffffttLLLLfLLLftttft111tt111i,,,,iiiii1tt111t1111tffttfftt11tfft
ttttttttttttfLLLffLLftttt1tttt11tti,,,iiiii111t11111111tffttfffttttttt
ffffffffttfftfLLfffLftttttffft11t1i,,,,,iii11ttttt11111ttt1tffttt11ttt
ffffffffttffftffftftttttffft111t11i,,,,iii1ii11tttt1111tt11ffttfftttff
fffffffttttttttttffftttttttt11111iii,,iiii1,,:,1tttt11111t11t1tfft1ttf
ffffffttttttttt1tffftttt11i,:,:i,,i,,,,,i11,,..,,:,i111tttt1111tt11ttf
fttttttttffttt11ttff1ii,:,....,,::,i,,i11ti,,.......,:,ittt111ttt1111t
tttttttfffftttt11tff,,,........,,,,iii1tti,,,,:,,::,....,i11111ttttt11
fffft1tffffttfft1ttt:..........:,,,,,,,i,...,,i,iii:.....,i1tt1ttffff1
fffft11tffftffftttti,..........,:,,,,,,,,...,,,iiii,......,tttt1tffftt
ffftt111ffttft11tt1,,...........,i,,,,,:,.....:,iii,......,tttt1tfft11
ftt1tttt1tttt1tttti,............,1,::,,,......,,iii,......,t11t11ttttt
111ttfftt111ttfffti.............,i,::::,.......,,,,,......,ttt11111tff
11ttttttt11tffffffi,............,::,:::,..........,.......:ttt1111tfff
11tttttttttttffffti,............,:::::,..............,....,ittt11ttttt
11tffttff11t1tfff1,...... ....,.,:::::,.....................,11111tfff
1tttttttfttttttti:,......:iiii,,,:::,:.......................,,111tfff
1ttftttttttfftti,,::,.,:,1i,,,,.,::::,.........................,i1tfff
11ttttttt1ttfft1:::,,,:,:,,,,i:.,:::,,..........................,1tfff
1ttttttttttfttt1,,...,...:,,,:, ,:::,..........................,11tttt
1tfttttfftttttfti,..,....,,,....,:::,.........................:111tttt
1tfffffffttttffft1,::::,,.......,:,:,..................     .:1111tttt
1ttttttttttttttttttt11111:......,:,,,..................,::,:1tt1111ttt
11t1111tttttttttttt111111:......,,.....................,1tttttt1111111
1111111ttttttttttttt11111:......,:,,....................:1t1ttttt11111
1111111ttttttttttttt1111t:......,,,::....................:1ttttttt1111
111111111ttttttttt1111111:......,::::,...................,itttttt11111
111111111111tttt111111111:......,,,::,....................:1ttttttt1tt`,

	`tfffLfttffffffLLffffttttttttttt1ttti,,,,::,i1tttt111111tffffffffftttff
fttffftffffffffffttfftttttt1ttt111i,,,,,,,,,,11111t1111111tffffLLffttf
LfttttffLLLfftttttfLftttttttffft1i::,,,,,,,,,:1t11t111tt1111ttfffffttt
LLffttfLLLfttfLfftfLLtttttfffffti:,,iiiiiiii::1tt11111tftt1111ttfffftt
LLLLftffftttffLffftfftttttffffft::,,iii1i111i:1ttt1111ttt1tttt11tffLft
LLLfttttffftffLfftfftttttfffffft,,,,,,,iiiiii,tttt11111111ttftttttttt1
LLLfttfftfLfffffffLLftttttfffft1ii,,,,,iiiiiitfttt11111ttt1ttt11tttt11
LLfftfLLftfffLLffLLLftttttttfftiii,,i,,ii111itfft111111tfft11t11tt1ttt
fttttffffftfLLLLfLLLftttft111tt1iii,,,,iiiii1tt111tt111tffttfftt11tfft
ttttttttttttfLLLffLLftttt1tttt11t1i,,,iiii1111t11111111tffttfffttttttt
ffffffffttftttfLfffLftttttfffft11i,,,,iiii111tfttt11111ttt1fffttt1ttt1
ffffffffttfftttfttttttttffft11111i,,,,,iii1tt11tttt1111t111ffttffttfff
fffffffttttttttttffftt11ttt1111t1i,,,iii1i1itt11ttt111111t11t1ttft1ttf
fffffftttttt1111tffftttt11i,::i1iii,,,iii1i,:,i11t11111tttt1111tt11ttf
fttttttttffttt11ttft1ii,:,....,i,,,ii,,i1t:,...,:,i1111tttt111ttt1111t
tttttttfffftttt11t1:,,........ii:,,iii1tt1,,,......,,,i1tt11111ttttt11
fffft1tffffttfft11,,..........:,::,ii11t1:,,,...,,:,...,it11tt11tfffft
fffft11tfffttfttt1,....,,......,:,,,,,,i,.......:iii,,,.,1t1ttt1tffftt
ffft1111tfttft11ti,.............,i,,,,,:........:iiiii:..it11tt1tfftt1
ttt1tttt1tttt1tt,,..............i1,,,,,,........,,iii,,..it111t11ttttt
111ttfft11111tft:...............,i,,,,:,.........:,ii:..,ittt111111tff
11ttttttt11ttff1,...............,::::::......... ,,ii:...itttt111ttfff
11ttttttt111tft,,...........  ..,:,:::,...........,:,:...,ttttt11ttfff
11ttttttf11111,:,..... .....,:::::::::...............,...:1ttt1111ffff
11ttttttt11ti,,...........,:,i,,,:::::..................,.:tt11111tfff
11ttttttt111,.,,,,...,,..,,,,,,i,:,::,....................,,i1t111tfff
11ttttttt11t:.......,....,:,,,,,:::,:,.......................:1t11tfff
1tttttttt11t1,............,,,...,::::.........................:111tttt
1ttttttfft1t1:,............  .. ,:::,........................,i111tttt
1tfftttfftttttt1,::,,::.........,:,,,........................i1111tttt
1tttttttttttttttttt111i.........,::,,.................   . .,t11111ttt
1111111tttttttttttt111i,........,:,,...................,..,itttt111111
1111111tttttttttttt111i,....... ,:....................i111ttttttt11111
1111111ttttttttttt1111i,....... ,:,::,................,i1t11tttttt1t11
111111111tttttttt11111i,........,::,:,..................,1ttttttt11111
1111111111tttt11111111i........ ,,:::,..................,1ttttttttttt1`,

	`tfffffttffffffLLffffttt1ttttttt1tt111ii11tft1ttttt11111tfftfffffftttff
fttfffttfffffffffttffftttt111ttt1:,,,,,,,:i1111111t111111ttffffLLffttf
LfttttffLLLfftttttfLLfttttttftt1,:,,,,,,,,,,1tt111t1111t111tttfffffttt
LLffttfLLLfttfffftfLLfttttffft1i,,,,,,,,,,::1tttt11111ttt11111ttfffftt
LLLLftffftttffLLfftfftttttffffi:,iii111111i:ittttt1111ttt1tttt11tfLLft
LLLfttttffftffLfftfftttttfffff,:,,,,iiiiiii:ittttt11111111ttttttttttt1
LLLfttfftfLftfffffLLftttttffff1i,,,,,ii,iiiittttft11111ttt1ttt1ttttt11
LLfftfLLftfffLfffLLLftttttttf1i1i,i,,iiii1i1tffft111111tfft11t11tt1ttt
fttttfffffttLLLLfLLLftttft1111iii,,,,ii1i1i11tt111t1111tffttfftt11tfft
ttttttttttttfLLLffLLfttttttttt11i,,,,iiii11t11111111111tffttfffttttttt
ffffffffttftttfLfffLftttttffft11i,,,iiiii1t11ttttt11111ttt1fffttt1tttt
ffffffffttffttfffttttttttfft1111i,,,,iii1111111tttt1111tt11ffttffttfff
fffffffttttttttttffftt11tttt11t1ii,,,iii11tfft11ttt111111t11tttfft1ttf
fffffftt1ttt1111fffft1tt11iiii1iiii,iiii1i:1tft11111111tttt111ttt11ttf
fttttt1ttfft11111tft11i,:,..,,ti,,,i,iii1,,,,:,i1111111tttt111ttt1111t
tttttttfffftttt11i,::,,......iti::,i,iitt:,,....,:,i111ttt111t1ttttt11
fffftttffffttft11:...........,ii::,i11tti,,,,.......,:itt111tt1ttffff1
ftfft11tfffttftt,,............,:,,,,,11,,..............,1111ttt1tffftt
ffft1111tfttft11:.............,:ii,,,,,,...............:11t11tt1tfft11
ttt1ttt11t1111t,,..............:i1,,,,:,...............,111111t11ftttt
111tttft11111t1,...............:ii,,,,,............,,,.,1tttt11111ttff
11ttttttt11ttfi................:,::,::,...........,,i,,:1tttttt11ttfff
11ttttttt11ttfi................,,,:::,............,,iiii1tttttft1ttfff
11ttttttt11111,........  .......:,:::,............:iiiii11ttftt11tffff
11ttttttt11ti:,....... ,:::,,:,.,,:::,............:,,ii:i11tt1111tffff
11ttttttt11i,.........,iii,,,,..,:::,.............::,,,.it1111tt11tfff
11ttttttt11,.,,,,..,,,:,,,,,,,..,:::,.............,::,..:i111ttt11tfff
11ttttttt11,,,,...... .:,,,,:...,:,,,.....................:111tt11tttt
1tttttttt111,..........,,,,.....,:,,.......................:111111tttt
1ttttttfft11i:,.....,....  .....,::,......................,1t11111tttt
1ttttttttttttt1ii,,,:...........,:,,.....................,it111111tttt
1111111tt1tttttttt11:...........,:,,..............      .,t11111111111
1111111ttttttttttt11,.......... ,,................,:,,,:i1tttt1t111111
1111111tttttttttt11i,...........,,,,..............:1t11ttttttttttt1111
111111111tttttt1111,........... ,,:,,..............:11ttttttttttttt111
11111111111tttt1111,........... ,,:,:...............:111tttttttttt11t1`,

	`tfffLfttffffffLLffffttt1tttttt11tffffffffffttttttt11111tffffffffftttff
fttfffttfffffffffttffftttt11111111i,i11ttt11tt1111t111111ttffffLLffttf
LfttttffLLLfftttttfLLfttttttfti:::,,,,:,i1ttfttt11t111tt111tttfffffttt
LLffttfLLLfttfffftfLLftttttft1,,,,,,,,,,,,tffffft11111tft11111ttffLftt
LLLLfttfttttfLLfLftffttttfffti,:,,,,iii,::1fffffttt111ttt1tttt1ttfLLft
LLLfttttffffffLfftfftttttfffi,,iiiii1111i:iffffftt11111111ttttttttttt1
LLLftffftfLftffffLLLftttttffi:,,,,iiiiiii:1ftttfftt1111ttt1ttt1ttttt11
LLfftfLLftftfLfffLLLLttttt1t1i,,,,,ii,iiii1ttffft111111tfft1tt11tt1tt1
fttttfffffttLLLLfLLLftttft11i1i,,,,i111iii111tt111t1111fffttfftt11tfft
ttttttttttttfLLLffLLftttt1tti1i,,,,ii1iii1tt11111111111tffttfffttttttt
ffffffffttftttfLfffLftttttfftt1,,,,iiii11ttt1ttftt11111tfttffftt111ttt
ffffffffttffttfftttttttttfft111i,,,iiii1t111t1ttttt1111tt11ffttffttfff
fffffffttttttttttfffttt1tft1111i,,,iiii111tfft11ttt111111t11tttfft1tff
fffffftt1ttt1111tffft11t111111iiii,,iii111fffft11111111ttft111ttt11tff
fttttt1ttfft1t11ttfftttt1iiii1i,,,iii1i1i,,i1ttt1111111tttt111tttt111t
tttt11tfffftttt11tff11i,:,,,:11i,,iii,it,....,:,i111111ttt111t1ttftt11
fffftttffffttfft1ii,:,,.....:11,:,,,,1tt,,,,.....,,,i11tt111tt1ttfffft
ftfft11tfffttft1,,..........:1ii,,ii1tti,,,..........,it1111ttt1tffftt
ffft1111tfttft1i,............,:,ii,,i1i,..............:111111tt1tfft11
ttt1ttt11t1111t,..............,,1i,,,,,...............,1111111t11ftttt
111ttttt11111tt:............. ,,ii,,,:,...............,i111tt11111tfff
11ttttttt11ttft,...............,,,,,:,.................ittttttt11ttfff
11ttttttt11ttf1,...............:,,,::,.......,,,.......,tttttfftttffff
11ttttttt1111t1,...............:,,,::........,,i,:,....,t1ttfttt1tffff
11ttttttt11t111,...............,::::,........:iiiii:...,t11tt1111tffff
11ttttttt11ttt1,..... ..........::::,.......,,iiii,,...it11111tt1tfffL
11ttttttt11tt1:...  .,..........,:::.........:,,i,,...,1t1t11tft11tfff
11ttttttt1tt1:,...:,ii:.........,::,.........,:,,:.....,it1111tt11tttt
1ttttttttt1t,::,,iiii,,.........,:,,...........,,,......,,11111111tttt
1ttttttfft1ti:,,,,,:,,,.........,:,,.......... ..........:tttt1111tttt
1ttttttttttt1,.,::::,,:.........,:,,...............,....,ttttt1111tttt
1111111111ttti:,,,::,:..........,:,,............ ......,ttttt111111111
1111111tttttttt1:,,,............,:,.............,.  .,,11tttttt1111111
1111111tttttttttt1:.............,,............ .11ii1ttttttt1111111111
111111111tttttt111,.............,::,,...........ittttttttttttt11111111
11111111111111111,..............,,,,:,..........:11tttttttttttttt11111`,

	`tfffLfttffffffLLffffttt1ttttttt1ttfffffffffttttttt1111ttfftffffLftttff
fttfffttffffffffftttfttttt111ttt1ttffffffttttt1111t111111ttffffLLffttf
LftttttfLLLfftttttffLftttttttt11ii11fffft1ttfftt11t111tt111ttfffLLfttt
LLffttfLLLfttfffftfLLfttttt1i:,:,,::,i1tttffffftt11111tft11111ttfLLftt
LLLLftffttttffLffftffttttfti:,,,,,,,,,:,tfffffttttt111tftttttt1ttffLft
LLLfttttffftffLfftffttttttti:::::,,,,,,,1ffttffttt11111111ttttttttttt1
LLLfftfftffftffffLLLfttttti,iiiii11111,:1ttfffffftt1111ttt1ttt1tftttt1
LLfftfLLftttfffffLLLLttttt,,i,,iiiiii1,:i11tfffft111111ffft1tt11tt1tt1
fttttfffffttLLLLfLLLftttft1ii,,,ii,,iii,tt11ttt111t1111tffttfftt11tfft
ttttttttttttfLLLffLLftttt11ti,,,iiii1iii1ttt1ttt1111111tfftffffttttttt
ffffffffttftttfLfffLfttttt11i,,,iii1iii1tttt1tfftt11111tfttffftt11tttt
ffffffffttfftttfftffttttfft11,,,iiiii11tt111ttttttt1111tt1tffttffttfff
fffffffttttttttttfffttt1tft11i,iii11i11111tfft11ttt11111tt11tttfft1tff
fffffftt1ttt111tfffftttt11111i,i,iiii111ttfffft11111111tttt11tttt11tff
fttttt1ttfftt111tfffttt111111,,i,,iii11i1ttffft11111111tttt111tttt11tt
tttt11tfffftttt11tffttt1t1ii1,,i,,iiii1,,::i11tt1tt1111ttt11tt1tfftt11
fffft1tffffttfft11tft11i::,,1ii,:,ii,it:,,...,,:,i11111tt111tt1ttfffft
ftfft11tfffttfttttt1i:,....,it1i,,,,i1t,,,,,.......,,111t111ttt1tffftt
ffft111ttfttft11i,:,........i1tt1ii1tt1:,,...........:1111111tt1tfftt1
ttt1ttt11t1111t,,...........,,,,11,,,,:..............,1t111111t1tftttt
111ttttt11111tt:..............:,ii,,:,................it111tt11111tfff
11ttttttt11ttf1,..............,,,,,::,................it1tttttt1ttffff
11ttttttt11ttfi................,,,:::,,:,:,...........,tttttttftttffff
11ttttttt1111ti................:,:::,.:,iii,,.........:111ttfftt1tffff
11ttttttt11t11,................,,:::,.,iiiii:.........:tt11ttt111tfffL
11ttttttt1tttt,................,::::,.:,,ii,,.........,ttt1111tt1tfffL
11tttt1tt11ttti.................:::,...,,,i,,.........itttt11ttt11ffff
11ttttttt11tt1,.................:::,....,:,:,.........,1tt1111tt11tttt
1ttttttttt1t11:.................:::,.... ..............:1111111t11tttt
1ttttttfft11ti,,................,::,.........,......... :ttttt1111tttt
1tttttttt1tt1iiii:..............::,,...................:1tttt111111ttt
11111111t111iiiiii:.............,:,,,.................:1ttttt111111111
11111111t11i,,iiii,.............,:,,,...........   .,,ttttttttt1111111
1111111tt1t1i,,ii,,.............,,........... :i,,i1tttttttt1111111111
1111111111ttti:,,...............,,,.......... :ttttttttttttttt1t111111
11111111111111,::...............:,,::,........,i11ttttttttttttttt11111`,

	`tfffffttffffffLLffffttttttttttt1tffffffffffttttttt11111tfffffffLftttff
fttfffttfffffffffttffttttt1111tt1ttffffffttttt1111t111111ttffffLLffttf
LfttttffLLLfftttttfLftttttttfffttttffffftttffftt11t111tt111ttfffLLfttt
LLffttfLLLfttfffftfLftttttt11i,,,,i1tftttffffffftt1111tftt1111tffLLftt
LLLLftffttttffLLfftffttttti:::,,,,::,i1tffffffffttt111tft1tttt1ttffLft
LLLfttttffftffLfftfftttt11,,,,,,,,:,,,,tffffffffttt1111t11ttttttttttt1
LLLftffftfLftffffLLLftttt1i,,,,iiiii:,:1tttfffffftt1111ttt1ttt1tfttt11
LLfftfLLftttfLfffLLLLttt1i1iiiii1111i::1111ttffft111111ffft1tt11tt1tt1
fttttfffffttLLLLfLLLftttt11,,,iiiiiii:,ttt111tt111tt111tffttfftt11tfft
ttttttttttttfLLLffLLftttt1i,,,iii,iii,ittfft11tt1111111tfftffffttttttt
ffffffffttfttfLLfffLfttt111,,,i111111iitfttt1tfftt11111tfttfffttt1tttt
ffffffffttffttffftfftttttt1i,,iiiii1i1tft11tttttttt1111tt1tffttffttfff
fffffffttttttttttfffttt1ttti,,,i111iitt111tfft11ttt111111t1ttttfft1tff
fffffftt1ttt111tfffftttt111ii,iii,ii11111tfffft11111111tfft11tttt11tff
fttttt1ttfft1111tffftttt1111,,,i,,,i111ttttffftt1111111tttt111tttt11tt
tttt1ttfffftttt11tfft111tt1i,,,i,,i1i1,i11tffftt11t1111ttt11tt1tfftt11
fffftttffffttfft11tft11ii,ii,,,,:,iii1,.,,,:,iii1ttt111tt111tt1ttfffft
ftfft11tfffttftttttt1,:,..:11,,i,,,,11:,,,.....,,:i1111tt111ttt1tffftt
ffft1111tfttft1111i:,.....,i11iiiii1tt:.,..........:111111t11tt1tffft1
ttt1ttt11t11111,:,.........,i11i1111t,..............,t11111111t1tftttt
111ttfft111111:.............,::,ii,,:...............:111111ttt111ttfff
11ttttttt11tt1,..............::,,,,,,...............,1111tttttt11tffLf
11ttttttt11tfi,..............,::,,::,,...............i111ttttfftttffff
11ttttttt111ti................::::::,:,,::,..........it111ttffttttfffL
11ttttttt1tt1,................,,::::.:iiiii,.........,11111ttt11ttfffL
11ttttttt11tti,................::::,.:ii,ii:.........,tttt1111ttttfffL
11ttttttt1ttf1.................,:::,.,:,iii:.. ......:ttttt11tft11tfff
11ttttttt1ttt,::...............,:::,. .:,ii:..........,ttt1111tt11tttt
1ttttttttt1t111i,,..............:::,....,::,...........,1111111t11tttt
1ttttttfftt1iiiiii:.............:::,....  ............. :ttttt1111tttt
11ttttttttt1i,iiii,,............,:,,,................. ,1tttt111111ttt
111111111111i,iii1i,............,::,,.................:1tttt1111111111
11111111111t1,,,,,,.............,:,,,..........     .:1tttttt111111111
1111111111t11,,.................,:,,..........,,:::,1ttttttt1111111111
11111111111111:..,..............,,............:ttttttttttttttttt111111
111111111111111ii,..............,::,...........i111ttttttttttt11111t11`,

	`tfffffttffffffLLffffttttttttttt1tffffffffffttttttt1111ttfffffffLftttff
fttfffttfffffffffttffftttt1111tt1ttffffffttttt1111t111111ttffffLLLfttf
LfttttffLLLfftttttfLftttttttfffft1ttfffftttffftt11t111tt111ttfffLLfttt
LLffttfLLLfttfffftfLftttttfttttttttttftttffffffftt1111tftt1111tffLLftt
LLLLftffttttffLffftffttttt1ii::::,111tttffffffffttt111tft1tttt1ttffLft
LLLfttttffffffLfftffttttt1:,,,,,,,,:,ittffffffffttt1111t11ttttttttttt1
LLLftffftfLftffffLLLftttti::,::::::,,,,ttttfffffftt1111tft1ttt1tffttt1
LLfftfLLftttfLfffLLLfftt11iiiii111i,,,,t111tfffft111111ffft1tt11ttttt1
fttttfffffttLLLLfLLLftttt1iiiii1111i:,iftt111ttt11tt111tffttfftt11tfft
ttttttttttttfLLLffLLftttti,,,ii,iiii::1ttttt11tt1111111tfftffffttttttt
ffffffffttfttfLLfffLfttt1i,,i1i,iiii:,tftttt1tfftt11111tfttfffttt1tttt
fffffffftfffttffftffttttti,,iii111i1iitft11tttttttt1111tt11ffttffttfff
fffffffttttttttttfffttttti,,,ii1111ii1t111tfft11ttt111111t11tttfft1tff
fffffftt1ttt111tfffftttt1i,,,iiiiii11111ttfffft11111111tfftt1tttt11tff
fttttt1ttfft1111ttffttt111i,,ii,,,itt11ttttfffft1111111tttt111tttt11tt
tttt11tfffftttt11tfft1t1t1i,,,ii,,it111ttttffftt11t1111ttt111t1tfftt11
fffftttffffttfft1ttft11ii1,,,,i,:,i1,,i11tttft111tt1111tt111tt1tffffft
ftfft11tfffttftttt1i,:,,:1i,,,,i,,,1,.,,,,,i11111111111tt111ttt1tffftt
ffft111tffttft1ii:,.....,11iiiiiii1t:,,,.....,:,i111111111t11tt1tffft1
ttt1ttt11t11i,:,........,i1tt1ii11tt:,,,,........:t11111111111t1tffttt
111ttfftt11i,............::,,iiiii1,,,............it1111111ttt111ttfff
11ttttttt11,...............:,,,,:,,...............:111111tttttttttffLf
11ttttttt11:...............,,,,,,,:,..............,111111ttttfftttffff
11ttttttt11,...............,:,,,,:,...............,1111111ttffttttfffL
11ttttttt11:......  ........:,,,,:,...............,11111111ttt1tttfffL
11ttttttt11:.......:,,,:....,,,::::.......,::,....,1tt11tt1111tt1tfffL
11ttttttt11:.....:,,,,,:.....:,:::,.......,:iii,,.,1t1t11tt11tft1tffff
11ttttttt1i,.....:,:,,,:.....,::::,........,,iii,..,t1111t1111tt11tttt
1tttttttt1i,.....:::,,,,......::::,.......,,,ii,:..,,1t11111111t11tttt
1ttttttttti....,,,,:,,,.......,:::,,.......,,,,,,....,it11tttt1111tttt
11ttttttt1:,,.,,,..,,,.........::::,........,:,,..... ,1tttttt11111ttt
111111111i,,..,,,.... .........::::,..................itttt11111111111
1111111111,...,,...............,::,,,...........,,...,tttttt1111111111
11111111111,..... .............,::,,,.........   . .:1tttttt1111111111
1111111111ti,,,.,,..............::,,..........,,,::i1tttttttttt1111111
11111111111111i11,..............:,............,1tttttttttttttt11111111`,

	`tfffffttffffffLLffffttttttttttt1tffffffffffttttttt11111tfffffffLfttfff
fttfffttfffffffffttfftttttt11ttt1ttfffffftttttt11tt111111ttffffLLLfttf
LfttttffLLLfftttttfLftttttttfffft1tffffftttffftt11t111tt111ttfffLLfftt
LLffttfLLLfttfffftfLfttttttffffffttttftttffffffftt1111tftt1111tffLLftt
LLLLftffttttfLLffftfftttttftttttttffttttffffffffftt111tftttttttttfLLft
LLLfttttffffffffftffttttttiii,::::i1ttttffffffffftt1111t11tttttttttttt
LLLftffftfLfttfffLLLfttt11:,,,,,,,:::itttttfffffft11111tft1ttt1tffttt1
LLfftfLLftttfLfffLLLLfttt1i::::,,::,,,itt11tfffft111111ffft1tt11tt1tt1
fttttfffffttLLLLfLLLftttttiiiiii1i,:,,:ttt111ttt11tt111tffttfftt1ttfft
ttttttttttttfLLLffLLftttti,iiii111i,:,iftftt11ttt11111tffftffffttttttt
ffffffffttfttfLLfffLfttt1,,,ii,iiii,::1ffttt1tffft11111ffttfffttt11ttt
fffffffftfffttffftfftttt1,,i1i,,iii,:itft11tttttfttt111tt1tffttffttfff
fffffffttttttttttffftttt1,,iiii111iii1t111tfft1tttt111111t1ttttfftttff
fffffftt1ttt111ttffftttti,,iiiiiiiii1111ttfffft11111111tffttttttt11tff
fttttt1ttfft1111ttfftttt1,,,iiii,,itt11tffffffft1111111tttt111tftt11tt
tttt11tfffftttt11tfftttt1i,,,,ii,,it1111tftfffft11t1111ttt111t1tfftt11
fffftttffftttfft11tft111i,,,,,ii,,i111111tttft111tt1111ttt11tt1tffffft
ftfft11tfffttfttt11i,,:ii,,,,iiiii1,11tt111t11111tt1111tt111tft1tfffft
ffft1111tfttftii,:,....,1i,,,,ii111,,:,i1tt11tt11111111111t11tt1tffft1
tt11ttt1111i,:,.......,i1tiiii,,11i,,,..,:,i1t1t111t1111111111t1tffttt
111ttttt1i:...........,ii111iii11i:.,,......,:,1t111111111tttt111ttfLf
11tttt11t,..............,::,,i1ti:,............,1t1111111tttffttttffLf
11tttttt1,...............:,:,,,,,:..............,tt111111tttffftttffff
11tttttt1................,,,,,,,,:..............:t11111111ttffttttffLL
11tttttti............... .:,,,,,:,..............:1111111111tttttttffLL
11tttttt,............. .:,,,,,,:::..............:111tt11tt111ttt1tfffL
11ttttt1,..............:,i,,i,::::..............:1ttt1111tt11tff1tffff
11ttttt,..............:,,,,,i:::::..............,1111111111111tt11tttt
1tttttt:..............:,,,::,:::::..............,t1111111111111t11tttt
1ttttt1,........... .,,:,...,:::::..............,11i,,i111tttt1111tttt
11tttt,............,..... ...,::::...............:i,:,,i1ttttt11111ttt
111111,............,.........,::::,..............,:::,ii1tt11111111111
111111,............ .........,,,::,..............::::,i11tt11111111111
1111111:.......  .............,,:,,..............,::::i111t11111111111
11111111:....,,,,.............,:::,,..................,1t1111111111111
111111111iiii11t, .............,:,..,........       ..:1tttttt11111111`,

	`tfffffttffffffLLffffttt1ttttttt1tffffffffffttttttt11111tfffffffLffttff
fttfffttfffffffffttfftttttt11ttt1ttfffffftttttt11tt111111ttffffLLLfttf
LfttttffLLLfftttttfLftttttttfffft1ttfffffttffftt11t111tt111ttfffLLfttt
LLffttfLLLfttfffftfLfttttttffffffttttffttffffffftt1111tftt111ttffLLftt
LLLLfttfttttfLLffftfftttttfftffffffftttfffffffffftt111tft1tttttttfLLft
LLLfttttffffffffftfftttttt1iii::,i1ttttfffffffffftt1111t11tttttttttttt
LLLfftfftfLfttfffLLLfttt11:,::,,,,,:,1tttttfffffftt1111tft1ttt1tffttt1
LLfftfLLftttfLLffLLLffttt1,:,,,,,,,,,,ttt11tfffft111111ffft1tt11ttttt1
fttttfffffttLLLLfLLLftttt1iiiiiiii,:,,iftt111ttt11tt111tffttfftt11tfft
ttttttttttttfLLLffLLfttt1i,iii1111i:,:1ftfft11ttt111111tfftffffttttttt
ffffffffttffttfLfffLfttt1,,,i,iii1i:,,tffttt1tffft11111tfttfffttt11ttt
fffffffftffftttfffffttt11,,,ii,,iii::1fft11tttttfttt111tt1tffttffttfff
fffffffttttttttttfffttt1i,,iiiii11i,itt111tfft1tttt111111t1ttttfftttff
fffffftt1ttt111ttfffttt1i,,iiiiiiiii1111ttfffft11t11111tffttttttt11ttf
fttttt1ttfft1111ttffttt1i,,,iiii,,,1t11tffffffft1111111ttttt1ttfft11tt
tttt11tfffftttt11tfftt111,,iiiii,,,1111ttftfffft11t1111tttt1tt1tffttt1
ftfftttffftttfft11tft1i1i,,,,iii,,,111111tttft111tt1111ttt11tf1tffffft
ftttt11tfftttfttt11i,,,1,,,,,iiii,i1t11t111t11111tt1111tt111tft1ffffft
fttt1111tftttti,:,,...:1i,,,,iiii1i:,i1ttt111tt11111111111t11tt1tffft1
tt11ttt111i,:,,.......:11ii,,ii,11i,.,,:,11tt11tt11t1111111111t1tffttt
111ttttt1:............:1111iiii1ti,,,....,,,i1tttt11111111tttt111ttfLf
111ttt1t,...............,:,,i1tt1,,..........,itttt111111tttttttttffLf
11ttttt1,................:::::,i,:,...........,1ttt111111ttttfftttffff
11ttttti.................,,,,,,,,:,...........,1t111111111ttffttttffLL
11ttttti................,,,,,,,,::,............it1111111111tttttttffLL
11ttttt,.............:,iii,,:,,:::,............it1111111tt111tttttfffL
11tttti,............,,,iiii::,::::,...........,1ttttt111ttt11tft1tffff
11ttt1:.............:,,,iii,::::::............,1tttt1111111111tt1tttft
1tttti..............,,,,,:,,:::,,:............:tt11t11111111111t11tttt
1tttt:............. ,::,.  .::,,,:........... :t1111111t111ttt1111tttt
111ti.......................,:::::.............it111t11111tttt1111tttt
1111,.............,.........,:::::.............:1t11ii,:,1111111111111
1111:........................,,:::..............,ii,:,:,ii111111111111
11111:.......  ..............,,,:,...............,:::::,i1111111111111
111111:......,,..............,,,:,...............:::,::,i11t1111111111
1111111iiiiii1,...............,,,,.,..............,::::,11111111111111`,

	`tfffffttffffffLLLfffttt1ttttttt1tffffffffffttttttt11111tfffffffLffttff
fttfffttffffffffftttftttttt11ttt1ttfffffftttttt11tt111111ttffffLLLfttf
LfttttffLLLfftttttfffttttt1tfffft1ttfffffttfffft11t111tt111ttfffLLfttt
LLffttfLLLfttfffftfLftttttttfffffftttffttffffffftt1111tftt111ttffLLftt
LLLLfttfttttfLLffftffttttt111i,iitfftttfffffffffftt111tft1tttttttfLLft
LLLfttttffffffffftfftttt1,:,::,,,:,ittttffffffffftt1111t11tttttttttttt
LLLfttfftfffttfffLLLftt1i:,,,,,,,,,,:1tttttfffffftt1111tft1ttt1tffttt1
LLfftfLLftttfLLffLLLftt1i,,,,,,,i,:,:tftt11tfffft11111tffft1tt11ttttt1
fttttfffffttLLLLfLLLft1,iiiii11111i::tffft11tttt11t111tfffttfftt1ttfft
ttttttttttttfLLLffLLftt11,,,iiiiiii:,tffffft11tt1111111tfftffffttttttt
ffffffffttffttLLfffLft11i,,,,i,,iii,tffffttt1tffft11111tfttfffttt1tttt
fffffffftffffttfftfttt11i,,,i1iii1iitffft11tttttfft1111tt1tffttffttfff
fffffffttttttttttffftt11i,,,iii111ii1tt111tfft1ttft111111t1ttttfftttff
fffffftt1ttt111ttffftttti,,,,iiiiii11111ttfffft11t11111tfftttttttt1tff
fttttt1ttfft1111ttffttt1i,,,iiii,,itt11tffffffft1111111ttttt1ttfft11tt
tttt11tfffft1tt11tfft111,,,,,iii,,it111ttfffffft11t1111tttt1tt1tffftt1
ftfftttffftttfft11tti,,1,,,:,iii::i111111ttfft111tt1111ttt11tf1tffffft
ftttt11tfftttft1i,:,..,1i,,,,iii,i,:,,1t111t11111tt1111tt111tft1ffffft
fftt1111tt1ii:,,......,i1i,,,ii,11,...,:,i111tt11111111111t11tf1tffftt
tt11ttt11,,,..........,11t1iiii1t1:,,.....,:i111111t1111111111t1tffttt
111ttttti,.............:,,iii11t1i,,.........,ittt11111111tttt11tttfLf
111tttt1:................:,,,,,i,:,...........:ttt1111111tttffttttffLL
11ttttti.................,,,,ii,,:............,1ttt11111ttttffftttffff
11ttttti..................:,,,i,,:,...........,1tt11111111ttffttttffLL
11ttttt,..................,,,,,,::,...........,1t1111111111tttttttffLL
11ttttt,...........,,.....,,,:,:::............,111111111tt111tttttfffL
11tttti,........,,ii:......:,:::::............,ittttt111ttt11tft1tffff
11tttt,........:,,,,:,.....,,:::::.............ittttt1111t1111tt1ttftt
1tttt1,.......,:::,,,,......::::::.............,tt1t1111111ttt1tttfttt
1tttti.........,:::,:.......,:::::..........,,:,1t11111111ttff1111tttt
111t1,..........,:::........,:::::..........,:,,i1111t111ttttt11111ttt
1111,............,,.........,::::,.........,:,,,i111111111111111111111
11111:.......................,:::,.. ......,:,,ii111111111111111111111
111111:.......,..............,:::,..........,::i11t1111111111111111111
1111111:,,..,, ..............,,,::,............,1111111111111111111111
1111111111i11:...............,..,:,,.....    ..:1t11111111111111111111`,

	`tfffffttffffffLLffffttttttttttt1tffffffffffttttttt1111ttfffffffLfttfff
fttfffttffffffffftttfttttt1111tt1ttffffffttttt1111tt11111ttffffLLffttf
LfttttffLLLfftttttffftttt11tttffttttfffffttfffft11t111tt111ttfffLLfttt
LLffttfLLLfttfffftfLftt11ii,i,,ii1tttffttffffffftt1111tftt111ttffLLftt
LLLLfttfttttfLLffftfftt1:,,,,,,,,:iftttfffffffffftt111tft1tttttttfLLft
LLLfttttffffffffftff1i,,:,,,,,,,,:,ttttfffffffffftt1111t11tttttttttttt
LLLfttfftfffttfffLLLt,,,:,,,,,iiii:ittttttffffffftt1111tft1ttt1tftttt1
LLfftfLLftttfLLffLLLfi,:iiii111111,1ffftt11tfffft11111tffftt1t1tttttt1
fttttfffffttLLLLfffffi:,,,,,iiiiiiitffffft11tttt11t111tfffttffft1ttfft
ttttttttttttfLLLffLLti,,,,,i,,iiii1ftffffftt11tt1111111tfftffffttttttt
ffffffffttffttfLfffLtii,,,ii,,i1111tfffffttt1tffft11111tfttffftttttttt
fffffffftffffttfftftt1ii,,,,,,iii11ttffft11tttttfftt111tt1tffttffttfff
fffffffttttttttttffftt11i,,,,,i111111tt111tfft1ttft11111tt1ttttfftttff
fffffftt1ttt1111tfffttti,,,,,,iiiii11111ttfffft11t11111tfftttttttt1ttf
fttttt1ttfft11111tfft11i,,,:,,ii,,itt11ttfffffft1111111ttttt1ttfft11tt
tttt11ttffft1tt11ttti::i,,,:,iii,,,i,i1ttftfffft1tt1111tttt1tt1tffftt1
ftfft1tffftttff1i,:,...,1,,,,,i,:,i,..,:,i1tft111tt1111ttt11tf1tffffft
ftttt11tfft11i,:,......,11,iii,,iii,,,....,:,i1111t1111tt111tft1ffffft
fftt1111t1,,,..........,i1t11ii11i:.,........,111111111111t11tt1tffftt
tt11ttt11:..............,iiiiiii11,,,......,,,itt11t1111111111t1tffttt
111ttttt,................,,,,i,,ii,...........,ttt11111111ttt111tttfLf
111ttt11:.................:,,,i,::............,ttt1111111tttftttttffLL
11ttttt1,.................,,,,,,,:............:ttttt11111tttffftttffff
11ttttt1,..................:,,,:::............:1tt11111111ttffttttffLL
11ttttti...................,,,::::............,111111111111tttttttffLL
11ttttt:...................,::::::,:::,.......,1t1t11111tt111tttttfffL
11tttt1,....................,::::::,,,ii:......ittt11111ttt1ttff1tffff
11tttti.....................,:::::,,,,ii,,.....itttt1111tt111ttttttttt
1ttttt,........ .............:::::,,,,ii,,.....:1t1t1111111tttttttfttt
1ttttt:.... :i,:.............:::::..,,,i:.......it11111111ttff1111tftt
1111t1,...,,iiii,, ..........::::,.  .,:,......,111111111ttttt1111tttt
11111i..,.,,,,iii,...........:::::.......,....itt111111111111111111111
111111,....::,,ii,...........::::,..........,,111111111111111111111111
1111111, ...,:,,,,............,,:,.....,,:,11tt11111111111111111111111
1111111i:,....... ............,,:,.....:1tttttt11111111111111111111111
1111111111,.,. ..............,,,:,,....,it11ttt11111111111t11111111111`,

	`tfffffttffffffLLLfffttt1ttttttt1tffffffffffttttttt1111ttffffffLffttfff
ftttffttffffffffftttftt1t1iii11111tfffffftttttt11tt111111ttffffLLffttf
LfttttffLLLfftttttffft11i:,,,::::,1tfffftttfffft11t111tt111ttfffLLfttf
LLffttfLLLfttfffftfft1,:,,,,,,,,,:tttttttffffffftt1111tftt111tttfLLftt
LLLLfttfttttffLfftt1::,,,,,:::::,itfttttffffffffftt111tft1ttfttttfLLft
LLLfttttffftffffttf1,,,,:,iii1111itftttfffffffffftt1111t11ttfftttttttt
LLLfttfftfffttffffft:,,,:,iii111ii1tttttttffffffftt1111tft1ttt1tftttt1
LLfftfLLftttfLfffLft:,,:,,,,iiiiii1tffftt11tfffft111111ffftt1t1ttttttt
fttttfffffttfLLfffft,:::,,,iii,iii1tffffft11tttt11tt111fffttffft1ttfft
ttttttttttttfLLLffffti,,,i,,iii,iitfffffffft1tttt1t1111tfftffffttttttt
ffffffffttffttfLftfft1i,,,,,iiiii11ffffftttt1tffft11111tfttffftttttttt
ffffffftttfftttftttttt1,,,,,iiiii11ttffft11tttttfftt111tt1tffttffttfff
fffffffttttttttttfftt11i,,,,,,ii1111ttt111tfft1ttft11111tt1ttttfftttff
fffffft11ttt1111ttff1,i,,,,,,,ii,iiii111ttfffft11t11111tfftttttttt1tff
ftt1tt1tttft1111tt1,,.:i,,,,,iii::,,,,:,i1tfffft1111111ttttt1ttfft11tt
tttt11ttffft1tti,:,...,ii,,,,ii,::,,.....,:itftt1tt1111tttt1tt1tffftt1
ftfft1tttft1i,:,.......:11iiii,,,:,:,,,,,,..,t111tt1111ttt11tf1tffffft
ttttt11tti,:,..........,1tt1ii11i,,,,.,.,,,,,1111t11111tt111tft1tfffft
tttt1111,...............:11iiii,,1,.,.,,.,..,1t11111111111t11tf1tffftt
tt111tt1,................,,,,ii,i1:.........,111t1111111111111t1tffttt
1111ttt,..................:,,,,:,i:..........itttt11111111ttt111tttfLf
111ttt1:..................,,,,,:::,..,.......,tttt1111111tttffttttffLL
11tttt,....................:,,::::,,.........:ttttt111111ttttfftttffff
11tttt:....................,,:::,:,,,,:,.....:1ttt11111111ttffttttffLL
11ttt1,.....................:::,,::,iii,,.....itt1111111111tttttttffLL
11ttt,......................,::,,::,,,i,,.....:111111111tt111tttttfffL
11tt1:......................,:,..:,:,,i,.......:tttt1111ttt1ttffttffff
11tt1:......................,:,..:. .:,:......,itttt1111tt111ttttttftt
11tt1:...........  .........,,,.,:...........,1tt11tt111111ttt1ttttfff
1ttti,......... .,..........,,,.,:..........,it11111111111ttft1111tttt
11tt:..........,::,,........,,,.,:......,:i11t11111111111ttttt1111tttt
11111:........:,,ii,.........,,.,,.....i1ttt11t11t11111111111111111111
1111ti........,:,,i,,...........,,.....,t11111111111111111111111111111
11111ti,,......,:,,,. ......,:,.,:.....,1t1t11111111111111111111111111
11111111,........,,,........:,::,,......,t1tttt11111111111111111111111
1111111i.........   ........:,,::,......,1ttttt11111111111t11111111111`,

	`tfffffttfffffffffffft11111,:::,,i1tffffffffttttttt1111ttffffffLffttfff
ftttffttffffffffftttft1,::,,,,,,,,ifffffftttttt11tt111111ttffffLLffttf
LfttttffLLLfftttttff1,:,,,,,,,,,,:itfffftttfffft11t111tt111ttfffLLfttf
LLffttfLLLfttfffttt1:,,,,:::,,i,,1tttftttffffffftt1111tftt111tttfLLftt
LLLLfttfttttffLfftt,,,,,,:,ii11t11tfttttffffffffftt111tft1ttfttttfLLft
LLLfttttffffffffttt,,,,,:,,,iii1111ttttfffffffffftt1111t11ttffttttttt1
LLLfttfftfffttttfft:,,,,:,,,iiiiii1tttttttffffffftt1111tft1ttt1tftttt1
LLfftfffftttfLftfLfi,,,,,,,iii1iii1tffftt11tfffft111111ffftt1t1ttttttt
fttttfffffttfLLftfffi,,,i,,,,iiiii1fffffft11tttt11tt111fffttffft1ttfft
tttttttttttttffftffft,,,,,,,,iiii1tttfffffft1tttt1t1111tfftffffttttttt
ffffffffttftttffftfft1,,,,,,,,iii1ttfffffttt1tffft11111tfttfLftttttttt
ffffffftttfftttftttt1i,,,,,,,,ii111tffftt11tttttfftt111tt1tffttffttfff
ffffffft1tt1111ttt1::i,,,,,,,i1i,i111ttt11tfft1ttft11111tt1ttttfftttff
ffffftt111tt111i,:,..,i,,i,iii11,,,,,::,1tfffft11t11111tfftttttttt1tff
ttttt111ttt1,:,......,ii,,iii,11,:,,,...,:1fffft1111111ttftt1ttfft11tt
1tt111tt1,:,..........,11iii,1t1,:,,,,,,,.:tfttt1tt1111tttt1tt1tffftt1
tttt111:..............:1tt1ii11i::,:,,,,,,,if1111tt1111ttt11tf1tffffft
tttt11,................:ii,,i,:,,,,,.,,.,..it1111111111tt111tft1tfffft
tttt1i,............,....:,,,,,,:,1i,..,,,..,ttt11111111111t11tt1tffftt
tt111,..................,,,,,,:,:1i,.......it111t11t1111111111t1tffttt
11111,...................:,,,::,,,,,.......it1tttt11111111ttt111tttfLf
111t,....................,,,:::,.:,..,.....:tttttt1111111tttffttttffLL
11t1,.....................,,:::..:,........,1tttttt111111ttttfftttffff
11ti......................::::,..:,.....,:::i1ttt111111111ttffttttffLL
11t:......................,:::,..:,....,:,iiiitt11111111111tttttttffLL
111,.......................,:,...:,....:,,iiii1111111111tt111tftttfffL
11t,...................,,,,,,,...,......,,,ii1111ttt1111ttt1tfffttffff
11i,.................,::::::,....:........,,:i1111tt1111tt111ttttttfff
1t,.................,:,,,,,,,....:....  .. .,1t11ttt1111111ttt1ttttfff
111:................,::::,i:.....:.,..,,,:,i11tttt111111111tft1111tttt
11t1:.   ............::::,:.....,:....it11111tt1tt1111111ttttt1111tttt
11111i:,..............,,,,......,:....:1tt111111t111111111111111111111
11111t11:........,.... .  ..,....:.....it1t111t11111111111111111111111
11111111:,..................:....:.....:1ttt11t11t11111111111111111111
1111111i,...................:....:......it1ttttt1tt11111111111t1111111
1111111,....................,,..,:,.....:tttttttt111111111111111111111`,

	`tffffftttffffffffffttt11,::,,,,,:,1ffffffffttttttt1111ttfffffffLfttfff
ftttffttffffffffttttt1,:,,,,,,,,,,:tffffftttttt111t111111ttffffLLffttf
LftttttffLLfttttttft,,,,,,,,,:::,:itfffftttfffft11t111tt111ttfffLLfttf
LLffttfLLLfttfffttt,,,,,,,:,ii1111tttttttffffffftt1111tftt111tttfLLftt
LLLffttfttttffffft1:,,,,,:,iii11111fttttffffffffftt111tft1ttft11tffLft
LLffttttffftfffftt1:,,,,,:,,,iiiii1ftttfffffffffftt1111t11ttffttttttt1
LLLfttfttffftttttft::,,::,,iiiiiii1ttttttttfffffftt1111tft1ttt1tffttt1
LffftfffttttffftfLfi:,i,,,,,iiiiii1tffftt11tfffft111111ffftt1tt1tt1ttt
ftttttffffttfLLftfff,:,,,,,,,iiiii1tffffftt1tttt11tt111fffttffft1ttfft
tttttttttttttffffffft,,,,,,,,iiii1ttffffffft1ttt1111111tfftffffttttttt
ffffffffttft1tfffttt1i,,,,,,,,iii1ttfffftttt1tffft11111tfttffftttttttt
ffffffftttffttttti::i,,,i,,,,ii1111tttttt11tttttfftt111tt1tffttffttfff
ffffffft1ttt11i,:,.,ii,,iiiii11:i1111tt111tfft1ttft11111tt1ttttffttttf
ffffftt111i,:,......:11,,iiiii1::,,,::,i1ttffft11t11111tfftttttttt1ttf
tt1111ii,:,.........,it1iii,,1ti::,,,..,:1tfffft1111111ttttt11tfft11tt
11111,,..............,1111ii1ti:,:,,,,,,.:tffftt11t1111tttt1tt1tffftt1
tttti,................:,,,i,,,,,:,,:,,,,,,itt1111tt1111ttt11tf1tffffft
tttt:..........,,......:,,,,,,,.,i,,,,,,,.,t111111111111t111tft1tfffft
ttt1,............,.....,,,,,,:,.:1i,......,11t111111111111t11tt1tffftt
t11,....................:,,,,:,.:1i,......:tt11111111111111111t1tffttt
11i,....................,:,:,:.,,,:,......:111tttt11111111tttt11tttfLf
11,......................,,:,:..,:,.......:11tttttt111111tttffttttffLL
1i,.......................:::,..,:,.......,1tttttttt11111ttttfftttffff
1,........................,:::,..:,........,11tttt11111111ttffttttffLL
1,......................,:,,,,,,,:,........,111tt1111111111tttttttffLL
1i.....................:::,,,,,:.:,....... ,1t11111111111t111tftttfffL
1:....................,:::,,,,,.,:,..........,:,iii11111ttt11tffttffff
1,....................,:::,,,,,.,:..........   .::,,,,ii11111ttftttftf
1i................. ....,,,,.....:.... ........,:::,,,,ii111tt1tttffff
1ti,.     .........,...   .. ....,..:,..    ....,:::,,,iii1ttt1111tttt
11t1ii,,,,.......................:..,.:1,:,,..   ..,,i1i1t1ttt1111tttt
1111ttttt:,......................:..  ,tttt111i,,,,i1tt111111111111111
111111111:,..................,...:.....it1t1ttttttttt11111111111111111
111111111,,..................,...:.... ,tttt11t11t11111111111111111111
111111111:...................,..,:.....:ttttttttttt1111111111111111111
11111111i,...................:,.,:.....,1ttttttttt11111111111111111111`,

	`tffffftttffffffLffftt111i,,:::::,1tffffffffttttttt11111tfffffffLfftfff
ftttffttffffffffttttt1i,,,,,,,,,,,,fffffftttttt111t111111ttffffLLffttf
LftttttffLLftttttttti:,,,,,,,,,,,,,tfffftttfffft11t111tt111tttffLLfttf
LLffttfLLffttfffttti,,,,,,:,,,ii,itttttttffffffftt1111tftt1tttttfLLftt
LLLLfttfttttffffft1:,,,,::,i111tt1tfttttffffffffftt111tft1ttft11tffLft
LLLfttttffftfffftti,,,,,,:,,iiii1i1ftttfffffffffft11111t11ttfttttttttt
LLLfttfttffftttttf1,,::,:,iiiiiiii1ttttttttfffffftt1111tft1ttt1ttfttt1
fffftfffttttffftfLf,:ii,,,,iii11ii1tfffft11tfffft111111ffft11t1ttt1tt1
ttttttffffttfLLftff1:,ii,,,,iiiiii1ffffftt11tttt1ttt111fffttfffttttfft
tttttttttttttfffffff1,,,,,,,iiiii1tffffffftt1ttt1111111tfftffffttttttt
fffffffftttt1tffftft1i,,,,,,,iiii1ttfffffttt1tffft11111tfttfffttt1tttt
ffffffft1tftt1tft1,ii,,,,,,,,ii1111ttftft11tttttfft1111tt1tffttffttfff
ffffffft1tttt1ii,,.:i,,,i,iii111t111ttt111tffttttftt1111tt1ttttfft1tff
fftfttt111ii,:,....,i1,,iiiii1i:,,,i11111tfffft11t11111tfftttttttt1tff
tt1111ii,:,.........,11iiii,,11,::,,,:,1tftfffft1111111tttt11ttfft11tt
11111:,.............:1tt1iii1t1,,:,:,,.,:1tffftt11t1111ttt11tt1tffftt1
tttt,................:iiiiii11,:,:,:,,,,.:tttt111tt1111ttt11tf1tffffft
ttt1,.................,,,,,:,:.,:i:,,,,,,,it11111111111tt111tft1tfffft
ttti...................:,,,,,:,,:1,.......,t1tt11111111111t11tf1tffftt
tt1:...................,,:,,,:..,1,.......:tt111t11t1111111111t1tffttt
11,.....................:,:,,,..,,:.......,111tttt11111111tttt11tttfLf
11,.....................,::,,,..,:,.......,1ttttttt111111tttffttttffLL
1,.......................:::,,..,:,........itttttttt1111tttttfftttffff
1,.......................,,,,,:,,:,........,11tttt11111111ttffttttffLL
i,.....................,::,i,,,,,:.........:111tt1111111111tttttttffLL
i.....................,:::,i,i:.,:,........,1t1111t111111t111ttfttffLL
,.....................,:::,,,,,.,:..........1tt111tt11111ttt1tffttffff
:......................:::::,...,:..........:i1tttt1tt111t1111tttttftt
i.......................... ....,,.......... .,,:i11iii11111tt1t1ttfff
1,...     ......................,:............   .,:::,iiii1111111tftt
1t1i,,::::......................,:..,:..   .......,::,,,,,,iii11111ttt
11ttttttt1,.....................,:......:,..    ..,:,,,:,,,iiii1111111
111111111i,.................,::..,.... ,tt1ii,,,.. ..,:i,,iii111111111
111111111i,.................,,:,,,......1tttttt11i,,::itt11t1111111111
111111111i,................ ,,:::,......it1t1tttttttttt11tt11t11111111
111111111,................. ,,,::,......it1tttttttt1111111t11111111111`,

	`tffffftttffffffLffftt11111i,iiii1tfffffffffttttttt11111tfffffffLfttfff
ftttffttffffffffttttt11,::,,,,,,::1fffffftttttt111t111111ttffffLLffttf
LftttttffLLftttttttf1,:,,,,,,,,,,,itfffftttfffft11t111tt111tttffLLfttf
LLffttfLLffttfffttt1,,,,,,,::::::itttftttffffffftt1111tftt11t11tfLLftt
LLLLfttfttttffLfft1:,,,,,:,ii11111ffttttffffffffftt111tft1ttft11tffLft
LLLfttttffftfffftt1:,,,,:,,iii1111fftttfffffffffftt1111111ttfttttttttt
LLLfttfttffftttttf1,,,,,:,iiii,iii1ttttttttfffffftt1111ttt1ttt1ttfttt1
fffftfffttttffftfLt:,,,:,,iiiiiii11tfffft11tfffft111111ffft1tt1ttt1tt1
ttttttffftttfLLftff1,,,i,,,iii1iiitfffffft11tttt1ttt111fffttfffttttfft
tttttttttttttffftfffi,,,,,,iiiiii1tffffffftt1ttt1111111ffftffffttttttt
fffffffftttt1tffftffti,,,,,,iiiii1tttffffttt1tffft11111tfttfffttt1tttt
ffffffftttftt1tftt111i,,,,,,,,ii111tfftft11tttttfft1111tt1tffttffttfff
ffffffft1tt111111i::i,,,i,,,,i11t111ttt111tffttttftt1111tt1ttttfft1tff
fftfftt111111i,:,..,ii,,iiiii11:,,,111111tfffft11t11111tfftttttttt1tff
tt1111111i,:,.......:11,iiii,11,::,,,:i1tftfffft1111111tttt11ttfft11tt
11111i,:,...........:1t1iii,itt,,,,:.,,,:itffftt11t1111ttt11tt1tffftt1
tttt1:..............,i111ii1tti:,,,:,,,,.,1tft111tt1111ttt11tf1tffffft
tttt,.................:,,,,,,,,.:i:.,,,,,.,t11111111111tt111tft1tfffft
ttt1,..................:,,,,,:,.:1,.....,.:11tt11111111111t11tf1tffftt
tt11,..................,,,,,,:,.,1,.......,1t111t11t1111111111tttffttt
111,....................::,,,,..,,:.......,i11tttt11111111tttt11tttfLf
11i.....................,::,,,..,:.........ittttttt111111tttffttttffLL
11,......................::::,..,:,........itttttttt11111ttttfftttffff
1,.......................,:::,..,:,........,11tttt11111111ttffttttffLL
1:........................,:,,..,:,........,111tt1111111111tftttttffLL
1,.......................,,::,..,:..........it11111t11111t111tttttffLL
i......................::,,,,,:.,:..........,t11tttt11111tt11tffttffff
:....................,:::,i,,i,.,:..........,ttttttt1111111111tttttftt
,....................,:::,i,,:..,,...........:,iii1111111111tt1t1ttfff
1:...................,:::,,,,...,,........... .,:,,,:,i1111ttt1111tttt
11,,..    ...............,,,....,:....  .......,:::::,,11tt11111111tt1
11t11i,,,:.......,.....   ......,,....,,.   ...,::,::,i111111111111111
11111tttti,...................,.,,... :1i:,.    .,,,,i1111111111111111
111111111i,...................:,,,.....ittt1i:,.. ,1tttt11111111111111
111111111,,...................,,::.... .111tttt1ii1ttt11111111t1111111
111111111,....................,:::......it1t11ttttt1111111111111111111`,

	`tffffftttffffffLfffttt1111ttttt1tffffffffffttttttt11111tfffffffLfttfff
ftttftttffffffffttttft1i,:,,,,,i11tfffffftttttt111t111111ttffffLLffttf
LftttttffLLfttttttfft1,:,,,,,,,,:1ttfffftttfffft11t111tt111tttffLLfttf
LLffttfLLffttfffttti::,,,,,,,,,,,ttttftttffffffftt1111tftt11t11tfLLftt
LLLLfttfttttfffffti,,,,::,,,iiiii1ffttttffffffffftt111tft1ttft11tffLft
LLLftttttfftfffftt1:,,,,iii1111111fftttfffffffffftt1111111ttfttttttttt
LLLfttfttffftttttff:,,:,,,iiii1ii1tttttttttfffffftt1111ttt1ttt1ttfttt1
fffttfffttttffftfLt,:,:,,,,iiiiii11tfffft11tfffft11111tffft1tt11tt1ttt
ttttttffftttfLLftff1,,,,,iiiii,i11ttfffftt11tttt1ttt11tfffttffft1ttfft
ttttttttttttfffLtfffii,,,,,ii,,iittffffffftt1ttt1111111ffftffffttttttt
ffffffffttft1tffftffti,,,,,iiiii1ttttffffttt1tffft11111tfttfffttt1tttt
ffffffftttfttttfttttti,,,,,iiii1111tfftft11ttttffff1111tt1tffttffttfff
ffffffft1tt11111tft11i,,,,,,,iit11111tt111tffttttft11111tt1ttttfft1tff
fftfftt1111111111i:,,i,,,,,,ii1i,,,111111tfffft11t11111tfftttttttt1tff
tt11t111tttt1i,:,...:1i,ii,ii,1i:,,i,,i1tftfffft1111111tttt11ttfft11tt
111111ttt1,:,,.......,ti,iii,it1,,,,,..,:i1fffft11t1111ttt11tt1tffftt1
tttt111,:,...........,tt11iii1t1::,,,,,,,.,1ft111tt1111ttt11tf1tffffft
tttt11,..............:i11iii111:,,,,,,,,,,.,t1111t11111tt111tft1tfffft
ttttt,.................:,,,,,,:.,ii,....,,.:1tt11111111111t11tf1tffftt
tt111:.................,,,i,,,,.,1i,.......:t111111t1111111111tttffttt
11111,..................:,,,,,,,,,,,.......:11tttt11111111tttt11tttfLf
1111i...................,,,,,:,..:,........:1tttttt111111tttffttttfLLL
1111:....................:i,,:,.,:,........,1ttttttt1111ttttffftttffff
11ti,.....................,,::,..:,........,i1tttt11111111ttffttttffLL
11t,......................,::,...:,.........,11t11111111111ttt1tttffLL
11t:......................,::,...:,...,,....it11111111111t111tttttfffL
111,.......................,,,...:,...,:,,,,itt11ttt1111ttt1ttftttffff
11i.........................,,...:,...,:,iiii11111t111111t111ttttttttt
11,.........................,,...:....,:,,iii11111111111111ttttt1ttfff
11,...................,,,,..,....,......,:,ii1ttt1111111111ttt1111tftt
11t,...............,:::::::..,...:.........,,it11tt1111111tttt1111ttt1
1111,.............:::::,,i:......:.....     :1ttttt1111111111111111111
11111,,.. ........,::::,,:.......,... ,i,,,i1tttt111111111111111111111
11111t11i,.........,:::,,. ..,:,.:.....itttttttt1111111t11111111111111
11111111t,...........,,,.... ,:,,:......it1111ttt111111111111111111111
111111111,..........    .... .:,,:,. ...:tttt11tt111111111111111111111`,

	`tffffftttffffffLfffttttt11tttt11tffffffffffttttttt1111ttfffffffLfttfff
ftttffttffffffffftttft1tttt11ttt1ttfffffftttttt111t111111ttffffLLffttf
LftttttffLLftttttttfft11iiiiiittt1ttfffffttfffft11t111tt111tttffLLfttf
LLffttfLLffttfffttfft,::,,,,,,:itttttffttffffffftt1111tftt11t11tfLLftt
LLLLfttfttttffffft11i:,,,,,,,,:,,fffttttffffffffftt111tft1ttft11tffLft
LLLfttttfffttffftt1,:::::,,,iii,:tfftttfffffffffftt1111111ttfttttttttt
LLLfttfttfLfttttfft,:,iii111111i,1ttttttttffffffftt1111ttt1ttt1ttftt11
LffftfffttttfffffLf,:,,,,ii1iiiii11tfffft11tfffft11111tffft1tt11tt1ttt
ftttttffffttfLLffffi:i,,,,,ii,ii11ttffffft11tttt1ttt11tfffttffft1ttfft
ttttttttttttfLfftffi,ii,ii,i11111tfffffffftt1ttt1111111ffftffffttttttt
ffffffffttftttfLftf1,,,,,,,,i1111tttffffftttttffft11111tfttfffttt1tttt
ffffffftttfftttftttt1i,,,,,iiii1111ttftft11ttttffff1111tt1tffttffttfff
ffffffft1ttt1111tfftt1i,,,,iii1tt111ttt111tffttttft11111tt1ttttfft1tff
ffttftt111t11111ttfft1,,,,,,ii1ti,,i11111ttffft11t11111tfft1tttttt1tff
tt111111tttt11111tti:i,,,,,,iii1,:,i1i1ttftfffft1111111tttt11ttfft11tt
1t1111ttfftt1tt1i,:..,i,,,,,iii1i:,,:,,:,i1ffftt11t1111ttt11tt1tffftt1
ttttt1tttttt1i,,.....:11,ii,,,iti:,,,,,...,:1t111tt1111ttt11tf1tffffft
ttttt111tti:,....,...,i11iii,i1ti,i,,,,,,,,.,i111t11111tt111tft1tfffft
tttt1111:,...........,i1t1iii11i,,1,....,,...,t11111111111t11tf1tffftt
tt111tt,..............,:,i,,,,,,.i1,.........:11111t1111111111tttffttt
1111ttt:................,,,i,,:,,:,,.........:tttt11111111tttt11tttfLf
111ttti,................,:,,,,:,.,:..........,ttttt111111tttffttttfLLL
111ttt,..................,,,,::,.,,..........,1tttt111111tttffftttffff
11tttt:...................:,,::..::,,.........ittt11111111ttffttttffLL
11ttt1:...................,,::,,::,,,i,,......,t11111111111ttt1tttffLL
11ttt1,....................,::,,::,,iii,. ...,it11tt11111t111tttttfffL
11t1ti.....................,:,,,,::,iii,.....,it11tt1111ttt1ttftttffff
11t1ti......................,:,. ,,,:,i:.......,11tt11111t111ttttttttt
1tttt,......................,,,..,, .,,.......,,11111111111ttttt1ttfff
11ttt:.......................,,..,,.....,....:1ttt111111111ttt1111tftt
11tt1,.......................,,..:,.........:1t11111111111tttt1111ttt1
11111,.......    ............,,..:,.....,,ii1t1t1111111111111111111111
111111,.... .,::,................:....,11ttt11tt11t1111111111111111111
111111,.....,i,,,,...............:,....it1111ttt1tt1111t11111111111111
1111111i:.,iii,,i:.......... ,::,:......itttttttt111111111111111111111
111111111i,:,,,ii,,......... ,,,::,.....:1ttttttt11111111111t111111111`,

	`tffffftttffffffLffffttttttttttt1tffffffffffttttttt1111ttfffffffLfttfff
ftttffttffffffffftttftttttt11ttt1ttfffffftttttt111t111111ttffffLLffttf
LftttttfLLLftttttttfftttttttfffft1tffffftttfffft11t111tt111tttffLLfttf
LLffttfLLLfttfffftffft11i,ii1ttffttttffttffffffftt1111tftt11t11tfLLfft
LLLffttfttttffLfLftti,:,:,,,:,,itffftttfffffffffftt111tft1ttft11tffLft
LLLfttttffftffffftt1,,,,,,,,,::,:tfftttfffffffffftt1111111ttfttttttttt
LLLfttfftfLftttffff1,::::,,iii,,:1ttttttttffffffftt1111ttt1ttt1ttfttt1
fffftfffftttfLfffLfi,iiii11111i,:i1tfffft11tfffft11111tffft1tt11tt1ttt
fttttfffffttfLLffLf,,,,iiiiiiii,:1tfffffft11tttt1ttt11tfffttffft1ttfft
ttttttttttttffLLffLt,,:,,ii,,iii,ttffffffftt1ttt1111111ffftffffttttttt
ffffffffttftttfLftf11i,,,i1ii11ii1tfffffftttttffft11111tfttfffttt1tttt
ffffffftttfftttfttt11i,,,,i1111i111ttffft11ttttffft1111tt1tffttffttfff
ffffffftttttt111tfft11,,,,iiiii1t1111tt111tffttttft11111tt1ttttfft1tff
fffffft111tt1111ttfft1i,,,iiiii1,,,111111ttffft11t11111tfft1tttttt1tff
tttttt11ttttt111tttft1i,,,,iiiii,,,1t11ttftfffft1111111tttt11ttfft11tt
tttt11tffftt1tt11tff1ii,,,,,iiii:,,i,,i1tttffftt11t1111ttt11tt1tffftt1
tttft1tttttt1ttt11i,,:i,,,,,iii,,,,:...,,,:,i1111tt1111ttt11tt1tffffft
ftttt11tttt1tfti,,...,i1i,,ii,,i1i,,.........,i11111111tt111tft1tfffft
tttt1111ttt11,:,......,11iii,,1t11,......,.,,.,i1111111111t11tt1tffftt
tt11ttt11i,:,.........:1tt1i,i1111,............,t11t1111111111tttffttt
1111tttt1,.............:,,,,,,,,:,:............:tt11111111tttt11tttfff
111tt11ti................:,,,,,:,:,............,1tt111111tttffttttfLLL
11tttt1ti................,,,,,::,:,.............ittt11111tttffftttffff
11tttttt,.................:,,:::,:,.............:tt1111111ttffttttffLL
11tttttt,.................,:,:::,:.,:,,.........,it11111111ttt1tttffLL
11tttttt:..................,::::,:,,,,,,,:,......itt11111t111tttttffLL
111t1ttt,...................,::,,:,:,,,iii,.....,1tt1111ttt11tftttffff
11tt1ttt:...................,::,,:::,,iiii,,.....:i1t1111t111ttttttttt
1ttttttt:....................::,,:..,,:,ii:...... .1t111111ttttt1ttfff
11tttttt:....................,::,:.   ..::,......,1t1111111ttt1111tftt
11ttt1tt:.... ...............,:,,:..............,1t1111111tttt1111ttt1
11111111,,,,:,...............,:,,:....... ...,,:1ttt111111111111111111
11111111iiiiii:...............,,,:........,:i11tt1tt111111111111111111
11111111iiiiiii...............,,,:.......,1tttt11tt1111t11111111111111
1111111i,iiiiii,..............,.,:........ittt11t111111111111111111111
111111111iiii,,...............::::,,......:1ttttt11t111111111111111111`,

	`tfffffttfffffffLffffttttttttttt1tffffffffffttttttt1111ttffffffLLfttfff
ftttftttffffffffftttfttttt111ttt1ttfffffftttttt111tt11111ttffffLLffttf
LftttttfLLLftttttttffttttttttffft1tffffftttfffft11tt11tt111tttffLLfttf
LLffttfLLffttfffftfffttttttfffffftfttffttffffffftt1111tftt11t11tfLLfft
LLLffttfttttffffffttft111i,ii1ttfffftttffffffffffft111tft1ttft11tffLft
LLLfttttffffffffftft1i:::,,,,::,,1fftttfffffffffftt1111111ttfttttttttt
LLLfttfftfLffftfffLf1,,,,,,,,::,,:1tttttttffffffftt1111ttt1ttt1ttfttt1
fffftfLfftttfffffLLf1,,,,,iiiii:,:1tfffft11tfffft11111tffft1tt11tt1ttt
fttttfffffttfLLLffLfiiiii111111i::1tffffft11tttt1ttt11tfffttffft1ttfft
ttttttttttttfffLffLfii,,,iiiiiii,:1ftffffftt1ttt1111111ffftffffttttttt
ffffffffttfttffLfffLti,:,,ii,ii1,,ttffffftttttffft11111tfttffftttttttt
ffffffftttfftttftttt1i,,,ii11111ii1ttffft11ttttffft1111tt1tffttffttfff
ffffffftttttt111tfff1ii,,,iiiiiii1111tt111tffttttft11111tt1ttttfft1tff
fffffftt1ttt1111tffft1i,,,iiiiii,,i111111tfffft11t11111tfft1tttttt1tff
tttttt1ttfft1111tttftt1,,,iiiiii,:,tt11ttftfffft1111111ttttt1ttfft11tt
tttt11tfffft1tt11ttft11,,,,,iiii,,,1t11ttttffftt11t1111tftt1tt1tffftt1
tttft1ttfffttfft11tftiii,,,,iiii,,,,:,ii1tttft111tt1111ttt11tt1tffffft
ttttt11tfftttfttt11i:,ii,,,,,iii,i:.....,,::,i111t11111tt111tft1ffffft
fttt1111tft1ft1i,:,...,1i,,,,,,i11:...........,,1111111111t11tt1tffftt
ft11ttt111111,:,......,111ii,,i1t1:........,,...,11t1111111111tttffttt
111ttttt11,:,.........,i111i,i1t1,,.............,1t1111111tttt11t1tfLf
111ttt11ti,............,::::,,,i,:,..............it111111tttffttttfLLL
11ttttttti...............,,,,,,,,:...............:t111111tttffftttffff
11ttttttt,...............,:,,,,:::...............,1t111111ttffttttffLL
11ttttttt:................,,,,,:::................it1111111ttt1tttffLL
11ttttttt:................,:,,::::................:t1111tt111tttttffLL
11111tt1t,.................,::::::................,1t111ttt11tftttffff
11ttttttt:.,:,:............,::::::................,itt111t111ttt1ttttt
1tttttttt,:,iii:............,:::::...........,:,....:1t11111tttt11tfff
11ttttttt,,iiiii:. ..........:::::...........,,,,,:..it1111tft1111tttt
11ttttttt:,iiiii,............,::::,........,::,,iii,it111ttttt11111tt1
111111111:,,iiii,,...........,:,::,.........,,,iiii,it1111111111111111
111111111,:,:,,,,............,::::,...........:,,ii,it1111111111111111
111111111,::.... .............,:::,,.........  .,,:i1t1111111111111111
111111111i,,..................,,,:,,....... .,,..,,1111111111111111111
1111111111i,:,i,,..............,,:,,,........1tii11t111111111111111111`,

	`tfffffttfffffffLfffftttttttttt11tffffffffffttttttt1111ttfffffffLftttff
ftttftttffffffffftttfttttt111ttt1ttfffffftttttt111tt11111ttffffLLffttf
LftttttfLLLftttttttffttttttttffft1ttfffffttfffft11tt11tt111tttffLLfttf
LLffttfLLLfttfffftffftttttfffffffffttffttffffffftt1111tftt11t11tfLLfft
LLLffttfttttffLffftfftttttfftffffffftttffffffffffft111tft1ttft11tffLft
LLLfttttffffffffftfftt11iii,:,,i1tfftttfffffffffftt1111111ttfttttttttt
LLLfttfftfLffttfffLLft1,:::,,,,::,,1ttttttffffffftt1111ttt1ttt1ttfttt1
LLfftfLfftttfLfffLLLfti,,,,,:::::,,,fffft11tfffft11111tffft1tt11tt1ttt
fttttfffffttLLLLffLLf1ii,,iiii11i:,,tfffft11tttt1ttt11tfffttffft1ttfft
ttttttttttttfffLffLLf1iiiiiii1111,:,tffffftt1ttt111111tffftffffttttttt
ffffffffttfttffLfffLf111,,,ii,,i1i:itfffftttttffft11111tfttffftttttttt
fffffffttffftttftttttt1i,,,iii,iii,ifftft11tttttffft111tt1tffttffttfff
ffffffftttttt111tffft11i,,,ii11111111tt111tffftttft11111tt1ttttfft1tff
fffffftt1ttt1111fffft111i,,,iiiiiii11111ttfffft11111111tfft1tttttt1tff
tttttt1ttftt1111ttffttt1i,,,iiii,,,1t11ttttfffft1111111ttttt1ttfft11tt
tttt1ttfffftttt11ttft1t1i,,iiiii,:,1t111tttffftt11t1111tftt1tt1tffftt1
ttfft1tffffttfft11tft111i,,,,,ii,:,i11111ttttt1111t1111ttt11tt1tffffft
ftttt11tfftttfttt1ttt1i1i,,,,iiiiii::,i111tt11111t11111tt111tft1ffffft
fttt1111tfttft11tti,::,i1,,,,,iii11:...,,::,i1tt1111111111t11tt1tffftt
tt11ttt1111tt11i,,,....,1i,,,,,,i11:.,.......,,:i11t1111111111tttffttt
111ttttt111i,:,........,1t1ii,,i1ii,.............,t1111111tttt11t1tfLf
111tttttt11:...........:iiiiii1tti,..............,1t11111tttttttttfLLL
11ttttttt1i,.............,:::,,ii,,...............,t11111ttttfftttffff
11ttttttt1,..............,:,,,,,,:,...............,1111111ttffttttffLL
11ttttttt1:.......... ....:,,,,,,:,................it111111tttttttffLL
11ttttttt1,........,,:,:,.,,,,,,:,,................,t1111t111tttttffff
111t1tttt1,.......,,,,,:,..:,,::::,................,1t11ttt11tftttffff
11ttttttt1:......,,,,,,,:,.:,,::::,................,it111t1111tt11tttt
1tttttttti,......:,:,,,:...,:,::::,.................:1t11111tt1t11tfff
1tttttttti,......:,:,,:.....,:::::,..................:11111ttt1111tttt
1111ttttt,...,,,.,:::. ......:::::,...................,111tttt11111ttt
111111111,,,,,,,.............,:::::,..................,1t1111111111111
11111111i,,..:,..............,,::::,,.................,11tt11111111111
111111111,...................,:,,::,,.................,,ii11t111111111
111111111i,..................,:::::,,................:::,,,i111t111111
1111111111i,,,,i1:............::,:,,,,...........  ..::,,,,i1111111111`,

	`tfffffttfffffffLfffftttttttttt11tffffffffffttttttt11111tffffffffftttff
ftttffttffffffffftttfttttt111ttt1ttfffffftttttt11tt111111ttffffLLffttf
LftttttfLLLfttttttffffttttttfffft1tffffffttfffft11t111tt111tttffLLfttf
LLffttfLLLfttfffftfffttttttfffffftfttfftffffffffttt111tftt11t11tfLLLft
LLLffttfttttffLLfftffttttffffffffffftttfffffffffftt111tft1ttft11tffLft
LLLfttttffffffLLftffttttttt11i,ii1tftttfffffffffftt1111111ttfttttttttt
LLLfttfftfffffffffLLfttt1i,::,,,,:,,1tttttffffffftt1111ttt1ttt1ttfttt1
LLfftfLfftttfLfffLLLLttt1,,,,,,,,::,:iftt11tfffft11111tffft1tt11tt1ttt
fttttfffffttLLLLfLLLftt1i,,,,,,iiii:,,fftt11tttt1ttt11tfffttffft1ttfft
ttttttttttttffLLffLLftt1,iiiii11111i::tfffft1ttt111111tffftffffttttttt
ffffffffttfttffLfffLftt1,i,,,iiiiiii,,tfftttttffft11111tfttffftttttttt
fffffffftffftttfttftttttii,,,,ii,,ii,iftt11tttttfft1111tt1tffttffttfff
ffffffftttttt11ttfffttt11ii,,ii1111ii11111tffftttftt1111tt1ttttfft1tff
fffffftt1ttt111tffffttt1i1i,,,iiiiii1111ttfffft11t11111tfft1tttttt1ttf
tttttt1ttfff1111ttffttttt1i,,,,i,:,it11ttftfffft1111111ttttt1ttfft11tt
tttt11tfffftttt11tfftt11tt1,,,ii,,,1t111tttffftt11t1111tttt1tt1tffftt1
ttfft1tffffttfftt1tft1111t1,,,,i,:,i11111tttft111tt1111ttt11tt1tffffft
ftfft11tfffttfttttttt11111i,,:,iiii1i,i1t1tt11111t11111tt111tft1ffffft
fttt1111tfttft11tt1t1i,,,1i,,,,ii1ii1,.,:,ii1tt11111111111t11tt1tffftt
tt11ttt11t1t11tt11i:,...,11,,,,,it111,,.....,:,i11tt1111111111tttffttt
111ttftt111111i,:,......:1t1i,,,,i1t1,..........,:i1111111tttt11t1tfLf
111ttttt111tt,..........:1111iii1,11:.............,it1111tttttt1ttfLLL
11ttttttt11t1,...........,,:::,,i,,:...............:t1111ttttfft1tffff
11ttttttt111i..............:,:::,,::...............,1t1111ttffttttffLL
11ttttttt111:..............,,,,,,,:,................,t11111ttt1t1tffLL
11ttttttt11i,..........,::::,,,,,,:,................,1t11t111tttttffff
11tt1tttt11i.........:,,,,:,,,::::::.................,t1ttt11tftttffff
11ttttttt111,.......,,,,,,:,:,::::::,................:tttt1111tt11ttft
1tttttttt11,,......,,,,,,,:::,::::::,................,itt111t11t11ttft
1tttttttt11,,......:,,,,i,...,:::::,,.................:tt1tttt1111tttt
1tttttttt1,,.,.....,:,:......,:::::,,..................it1tttt11111ttt
1111111111:,,,......,,. .....,,::::,,..................,1t111111111111
1111111111,,....,,...........,,,:::::,..................itt11111111111
1111111111,,.................,:::::::,,................:1tt11111111111
11111111111:....  ............,:::,:,,.................,1tt11111111111
11111111111ii,,,,,i, .........,,::,,,,,................,:,,1tt11111111`,

	`tfffffttfffffffLfffftttttttttt11tffffffffffttttttt11111tfffffffffttfff
ftttffttffffffffftttfttttt111ttt1ttfffffftttttt11tt111111ttffffLLffttf
LftttttfLLLfttttttffLftttttttffft1ttfffffttfffft11t111tt111tttffLLfttf
LLffttfLLLfttfffftfLfttttttffffffffftftttfffffffttt111tftt11t11tfLLfft
LLLffttfttttfLLLfftffttttffftt1iii11tttfffffffffftt111tft1ttft11tffLft
LLLfttttffffffLffffftttttfti,,::,,,:,i1tffffffffftt1111111ttfttttttttt
LLLfttfftfLfffffffLLfttttt1,,,,,,,,::,:ittffffffftt1111ttt1ttt1tftttt1
LLfftfLfftttfLfffLLLLfttt1i,:::,,iiii,,,t11tfffft11111tffft1tt1ttt1ttt
fttttfffffttLLLLfLLLffttf1,,iiii11111i::1t11tttt1ttt11tfffttfftttttfft
ttttttttttttfffLffLLftttti:,,,,iiiiiii,,tftt1ttt111111tffftffffttttttt
ffffffffttfttffLfffLfttttti,,,,,iii,i1iiftttttffft11111tfttffftttttttt
fffffffftfffftffffftttttft11i,,,iiii1iiit11tttttfft1111tt11ffttffttfff
ffffffftttttt1tttffftttttti1i,,,ii11iii111tffftttftt1111tt1ttttfft1tff
fffffftt1ttt1111tfffftttt111i,,iiiiii111ttfffft11t11111tfftttttttt1ttf
tttttt1ttfftttt1ttfftttttt111i,i,:,ii11ttftfffft1111111ttttt1ttfft11tt
tttt11tfffftttt11tfftt1ttt111,,i,,,iii1ttttffftt11t1111ttt11tt1tffttt1
tttft1tffffttffttttft111ttt11,,,,:,i1i1,,11tft111tt1111tt111tt1tffffft
ftfft11tfftttffttttft1111iii1,,,i,iiii1i,,,:,,i11tt1111tt111tft1ffffft
fttt1111tfttft11tt1t1i,:,,,,1i,,i1ii,iti,,.....,:,i1t11111t11tt1tffftt
ft11ttt11tttt1tt11i,:,.....,11i,i1i,1tt,,,.........,:i11111111t1tffttt
111ttttt11111tft,,.........,11111i1ttti,.............,1111tttt1111tfff
111tttttt11tfff1,...........,:,,,,,11i:...............itttttttt1ttfffL
11ttttttt11tfffi............ ,:,,,,,,,,...............,t1ttttfft1tffff
11tttttttttt1tf,..............:,,,,,,,,...............:111ttttttttfffL
11ttttttt11t111:..............:,,,,,::,...............,1t11ttt1t1tffLL
11ttttttt11ttt1:..............,,,,::,::................itt111tttttffff
11ttttttt1ttfft:..............,,,,,::::................:1tt11tft1tffff
1tttttttt1tttt1,.,::,..........:,:::,::,...............,1t1111tt11tttt
1ttttttttt1ttt1,,:,,:..........,::::::,,................it11111111ttft
1tttttttttttttiiiiii,,.........,:::::::,................:ttttt1111tttt
1tttttttt1ttt1,,,,ii,,.........,:::::::,.................ittttt1111ttt
11111111111tt1,,,,ii,...........::::::::................ :1t1111111111
11111111111tt1,::,,i:...........,:::::::,...............,itttt11111111
11111111tttttt,,,,::............,::::::,,...............ittttt11111111
1111111111tttt1:,,..............,:::::,,,,.......,:,::,:1t111111111111
11111111111111111i,,i:..........,:,,,.,,.,.......::,,i,it1ttt111111111`,
}
