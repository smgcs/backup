;author: stone
;date:2019.07.27


!`::NextWindow()

; ## Menu ##
!M::WinMinimize, A
!#M::WinMinimizeAll
!Q::Send !{F4}

; File
!T::Send ^t
+!T::Send +^t
!N::Send ^n
+!N::Send +^n
!D::Send ^d
!O::Send ^o
!W::Send ^w
!+W::Send #w
!S::Send ^s
!P::Send ^p
!+P::Send #p

; Edit
!Z::Send ^z
+!Z::Send ^y
!X::Send ^x
!C::Send ^c
!V::Send ^v
!A::Send ^a

;terminal jump
!U::Send ^u
!F::Send ^f
!E::Send ^e
!K::Send ^k
![::Send ^[
CapsLock & k::Up
CapsLock & j::Down
CapsLock & h::Left
CapsLock & l::Right


;google reflash
!R::Send ^r

NextWindow()
{
  WinGetClass, cur_class, A
  acitve_id := 0
  DetectHiddenText, On
  WinGet, id, list,,, Program Manager
  ; don't break the loop
  Loop, %id%
  {
    this_id := id%A_Index% 
    WinGetClass, this_class, ahk_id %this_id%
    if (this_class != cur_class) {
      continue
    }
    if (acitve_id = 0) {
      active_id := this_id
    }
  }
  if (active_id != 0) {
    WinActivate, ahk_id %active_id%
  }
}

