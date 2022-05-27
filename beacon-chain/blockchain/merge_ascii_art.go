package blockchain

var mergeAsciiArt = 
`                                                                                                                                                  	                                                                                                                                                      
		                                                                                                                                                      
		                                 +?$$$$$$?*;                         ;*?$$$$?*;                         +!$$$$$$?!;                                   
		                                 !##$???$@##$+                      !&#@$??$&#@*                      +@#&$????$##+                                   
		                                 !##;     +@#&;                    !##*     ;$##*                     @#$      ;&#+                                   
		                                 !##;      !##+                   ;##$        @#@                     $#&*      ++;                                   
		                                 !##;     ;@#&;                   *##*        ?##;                    ;$##&$!+;                                       
		                                 !##?!!!?$##$+                    !##+        !##+                      ;!$@##&$!;                                    
		                                 !##@@@@$$!+                      *##*        ?##;                          ;*?@#&!                                   
		                                 !##;                             ;##$        @#@                    ;?$;       ?##+                                  
		                                 !##;                              ?##!     ;?##+                    ;##+      ;$#&;                                  
		                                 !##;                               !&#&$??$&#@*                     ;&#&$$??$$&#@+                                   
		                                 +??;                                ;*?$$$$?+                        ;+!?$$$$$!+                                     
		                                                                                                ;;;;                                                  
		                                      ;+!??$$$?!*+;                                       ;*?$@&&&&&@@$!*;                                            
		                                   *?@############&$?+                                 ;!@###############&$!;                                         
		                                ;!@####&@$????$$@#####@!                             ;?&####$?*++;++*!?@####&?;                                       
		                               *@###&$*;         ;*$&###@*                          *&###@!;            ;!@###&!                                      
		                              !###&!;               ;?&###?                        *####!                  *&###?                                     
		                             !###@+                   ;$###$                      +###&+                    ;$###?                                    
		                            ;###&;                      $###?   ;;+*!??$$$$$$$$??!$###*                      ;@###*                                   
		                            !###!                        &###?$@&#####################@$?!+;                  +###$                                   
		                            $###+                     ;*?&#################################&$?*;               &##&;                                  
		                            $###+                  ;!$&########################################&$!;            &###;                                  
		                            *###?               ;!$################################################$*;        +###@                                   
		                            ;@###+            +$&####################################################&?;      $###!                                   
		                             +###&+         *$##########################################################$+  ;$###$;                                   
		                              *&###?;     +$##############################################################?*@###$;                                    
		                               +$###&?+ ;$#####################################################################?;                                     
		                                 *@####@&#####################################################################*                                       
		                                   +$&##################@?!*++*!$&###################&$?*++*!?$###############&*                                      
		                                     $###############&?+         ;!@###############@!;         ;!@##############!                                     
		                                   ;$##############&!;              *&###########&!;              !&#############!                                    
		                                   $##############@+    +@*          ;$#########$;          +@*    ;$#############!                                   
		                                  ?##############$;    *###*           $#######$;          +&##!     ?#############+                                  
		                                 +##############$     !#####!           $#####@;          *#####?     ?############@;                                 
		                                 @#############@;    !#######!          ;&####+          *#######?     $############?                                 
		                                +##############+    ?#########?          $###$          !#########$    +############&;                                
		                                $#############$   ;$###########?         !###?         !###########$;   $############!                                
		                                @#############*   !#############!        ?###$        *#############?   +############@                                
		                               ;&############&;    +?@#######&$+;       ;&####;       ;+?@#######&?+;    @############;                               
		                               ;#############@        +$&#&$*;         ;$#####@;          +$&#&$*;       $############+                               
		                               ;#############$    *+    ;+;   ;*;     *&#######&!     ;*;   ;+;    +*;   $############*                               
		                                &############@    ;$@!;     +$@!    ;?###########$;    *@$*     ;*$@+    $############*                               
		                                $############&;    ;$#&$*+!@##*    +@#############@+    +&#@?++$&#@;     @############+                               
		                                *#############*      $######&+    !#################?    +@######$;     +#############;                               
		                                 @############$       ?####@+   ;$###################$;   ;@####$       $############@                                
		                                 *#############*       !##@;   +@#####################&*   ;$##?       +#############!                                
		                                  $#############+       *$;   ?#########################?;   $!       +&############&;                                
		                                  ;&#############!          +$###########################@+          *&#############?                                 
		                                   +##############$*;     ;?###############################$+      *$##############@;                                 
		                                    *###############&$?!!$###################################$?!?$&################+                                  
		                                     *###################################&@$$$$@&#################################!                                   
		                                      +&##############################&?+;      ;+?&#############################?                                    
		                                       ;$############################@;            ;@###########################!                                     
		                                         ?###########################*              *##########################*                                      
		                                          +@#############&$!+$#######?              ?########$+!$&###########@+                                       
		                                            !&###########&;   $#######$+          +$########?   +&##########?;                                        
		                                             ;?###########&*   *@#######@$!;   ;$@########@*   *##########$+                                          
		                                               ;?&##########?;  ;*$&####&$*    ;!$&####&$*   ;$#########@*                                            
		                                                 ;!@#########@!;   ;++*+;   ;*;   ;+*++;   ;!&########$*                                              
		                                                    *$&########&$*;      ;*$&#&$*;      +*$&#######&?+                                                
		                                                      ;*$&#########&@$$@&#########&@$$@&########&$*;                                                  
		                                                         ;+?$&##############################&$?+;                                                     
		                                                             ;+!?$@&###################&$$!+;                                                         
		                                                                   ;++*!??$$$$$$$?!!*+;                                                               
		                                                                                                                                                      
		                  ;;;           ;;+*++;   ;;;++;;;++;;;  ;+++;;;++++; ;;;         ;;      ;;;      ;;;++;;;+++;;   ;;;+++++++;   ;+++++;              
		                 ;@#&+        +$&&@$@&#?  @#@@@&#&@@@#$ ;$@@@@#&@@@@! !#@        !#$     +&#&;     !#&@@@#&@@@&&;  !#&@@@@@@@*   $#&@@@&&$+           
		                 $#?#@;      ?#&!;   *#$  &&;;;$#!;;*#$   ;;;*#@;;;;   $#?      +#&;    ;@#?#$     !#!;;*#@;;;@#;  ?#$;;;;;;;    @#*   ;!&#?          
		                *#$ ?#$     *#&;     ;+;  ++   $#!  ;+;      *#$       ;&#+     $#*     ?#? $#!    ;+;  +#@   ++   ?#$           $#*     ;&#*         
		               ;&&;  @#*    $#$                $#!           *#$        !#@;   !#$     +#@  ;&#;        +#@        ?#@??????!    $#*      $#$         
		               $#!;;;*#&;   $#?                $#!           *#$         $#?  ;&&;    ;@#*;;;!#$        +#@        ?#@??????!    $#*      $#$         
		              !##@@@@@&#$   !#@;               $#!           *#$         ;&#+ $#*     ?#&@@@@@##!       +#@        ?#$           $#*      @#*         
		             ;&&+;;;;;;@#*  ;$#$+     @$       $#!           *#$          *#@!#?     *#@;;;;;;+##+      +#@        ?#$           $#*    +$#$          
		             $#*       +#&;  ;?&#@$$$$#$    +$$&#@$$;   ;$$$$@##$$$$*      $##@;    ;&#+       !#@;   $$@##$$*     ?#&$$$$$$$!   @#@$$$@&@!           
		            ;**         +*;     +*!?!*+;    ;*******    ;***********+      ;**;     ;*+         **;   *******+     ;*********+   +*!!!!*;             
		                                                                                                                                                      	                                                                                                                                                      
`