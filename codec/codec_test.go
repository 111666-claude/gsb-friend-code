package codec

import "testing"

func TestEncodeHasRightLength(t *testing.T) {
	if len(Encode("alice")) != Length {
		t.Fatal("邀请码长度应该是 6")
	}
}

func TestVerifyAcceptsEncodedCode(t *testing.T) {
	if err := Verify(Encode("alice")); err != nil {
		t.Fatalf("自己编出来的码应该能过校验，得到 %v", err)
	}
}

func TestVerifyRejectsBadLength(t *testing.T) {
	if err := Verify("ABC"); err != ErrLength {
		t.Fatalf("长度不对应该报 ErrLength，得到 %v", err)
	}
}

func TestVerifyRejectsUnknownChar(t *testing.T) {
	if err := Verify("ABC$EF"); err != ErrChar {
		t.Fatalf("非法字符应该报 ErrChar，得到 %v", err)
	}
}
