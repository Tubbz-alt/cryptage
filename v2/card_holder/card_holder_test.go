package card_holder_test

import (
	"github.com/sorribas/shamir3pass"
	. "github.com/whereswaldon/cryptage/v2/card_holder"
	. "github.com/whereswaldon/cryptage/v2/types"
	"math/big"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var Faces []CardFace = []CardFace{"ACE", "KING", "QUEEN"}
var EncryptedFaces []*big.Int = []*big.Int{big.NewInt(0), big.NewInt(1), big.NewInt(2)}

var _ = Describe("CardHolder", func() {
	Describe("Creating a CardHolder from scratch", func() {
		Context("When the key is nil", func() {
			It("Should return an error", func() {
				holder, err := NewHolder(nil, Faces)
				Expect(err).ToNot(BeNil())
				Expect(holder).To(BeNil())
			})
		})
		Context("When the faces are nil", func() {
			It("Should return an error", func() {
				key := shamir3pass.GenerateKey(1024)
				holder, err := NewHolder(&key, nil)
				Expect(err).ToNot(BeNil())
				Expect(holder).To(BeNil())
			})
		})
		Context("When the faces are empty", func() {
			It("Should return an error", func() {
				key := shamir3pass.GenerateKey(1024)
				faces := make([]CardFace, 0)
				holder, err := NewHolder(&key, faces)
				Expect(err).ToNot(BeNil())
				Expect(holder).To(BeNil())
			})
		})
		Context("When the arguments are valid", func() {
			It("Should return a CardHolder an a nil error", func() {
				key := shamir3pass.GenerateKey(1024)
				holder, err := NewHolder(&key, Faces)
				Expect(err).To(BeNil())
				Expect(holder).ToNot(BeNil())
			})
		})
	})
	Describe("Creating a CardHolder from encrypted cards", func() {
		Context("When the key is nil", func() {
			It("Should return an error", func() {
				holder, err := HolderFromEncrypted(nil, EncryptedFaces)
				Expect(err).ToNot(BeNil())
				Expect(holder).To(BeNil())
			})
		})
		Context("When the faces are nil", func() {
			It("Should return an error", func() {
				key := shamir3pass.GenerateKey(1024)
				holder, err := HolderFromEncrypted(&key, nil)
				Expect(err).ToNot(BeNil())
				Expect(holder).To(BeNil())
			})
		})
		Context("When the faces are empty", func() {
			It("Should return an error", func() {
				key := shamir3pass.GenerateKey(1024)
				faces := make([]*big.Int, 0)
				holder, err := HolderFromEncrypted(&key, faces)
				Expect(err).ToNot(BeNil())
				Expect(holder).To(BeNil())
			})
		})
		Context("When the arguments are valid", func() {
			It("Should return a CardHolder an a nil error", func() {
				key := shamir3pass.GenerateKey(1024)
				holder, err := HolderFromEncrypted(&key, EncryptedFaces)
				Expect(err).To(BeNil())
				Expect(holder).ToNot(BeNil())
			})
		})
	})
	Describe("Setting the both encrypted values for all cards", func() {
		Context("When the input slice is empty", func() {
			It("Should return an error", func() {
				key := shamir3pass.GenerateKey(1024)
				holder, err := NewHolder(&key, Faces)
				Expect(err).To(BeNil())
				Expect(holder).ToNot(BeNil())
				err = holder.SetBothEncrypted([]*big.Int{})
				Expect(err).ToNot(BeNil())
			})
		})
		Context("When the input slice is nil", func() {
			It("Should return an error", func() {
				key := shamir3pass.GenerateKey(1024)
				holder, err := NewHolder(&key, Faces)
				Expect(err).To(BeNil())
				Expect(holder).ToNot(BeNil())
				err = holder.SetBothEncrypted(nil)
				Expect(err).ToNot(BeNil())
			})
		})
		Context("When the input slice has a different number of values", func() {
			It("Should return an error", func() {
				key := shamir3pass.GenerateKey(1024)
				holder, err := NewHolder(&key, Faces)
				Expect(err).To(BeNil())
				Expect(holder).ToNot(BeNil())
				err = holder.SetBothEncrypted([]*big.Int{big.NewInt(0)})
				Expect(err).ToNot(BeNil())
			})
		})
		Context("When the arguments are valid", func() {
			It("Should not error", func() {
				key := shamir3pass.GenerateKey(1024)
				holder, err := NewHolder(&key, Faces)
				Expect(err).To(BeNil())
				Expect(holder).ToNot(BeNil())
				err = holder.SetBothEncrypted(EncryptedFaces)
				Expect(err).To(BeNil())
			})
		})
	})
	Describe("Getting the local player's encrypted values for all cards", func() {
		Context("When the holder has only opponent's encrypted cards", func() {
			It("Should return an error", func() {
				key := shamir3pass.GenerateKey(1024)
				holder, err := HolderFromEncrypted(&key, EncryptedFaces)
				if err != nil {
					Skip("Holder failed to construct")
				}
				faces, ok, err := holder.GetAllMine()
				Expect(faces).ToNot(BeNil())
				Expect(ok).To(BeFalse())
				Expect(err).To(BeNil())
			})
		})
		Context("When holder has the right faces", func() {
			It("Should return the faces", func() {
				key := shamir3pass.GenerateKey(1024)
				holder, err := NewHolder(&key, Faces)
				if err != nil {
					Skip("Holder failed to construct")
				}
				faces, ok, err := holder.GetAllMine()
				Expect(faces).ToNot(BeNil())
				Expect(faces).ToNot(BeEmpty())
				Expect(ok).To(BeTrue())
				Expect(err).To(BeNil())

			})
		})
	})
	Describe("Getting a card face", func() {
		var encryptedFaces []*big.Int
		var key shamir3pass.Key
		var holder CardHolder
		var err error
		BeforeEach(func() {
			key = shamir3pass.GenerateKey(1024)
			encryptedFaces = make([]*big.Int, len(Faces))
			holder, err = NewHolder(&key, Faces)
		})
		Context("When the index is out of bounds", func() {
			It("Should return an error", func() {
				Expect(err).To(BeNil())
				Expect(holder).ToNot(BeNil())
				err = holder.SetBothEncrypted([]*big.Int{})
				Expect(err).ToNot(BeNil())
			})
		})
		Context("When the card can't be decrypted", func() {
			It("Should return an error", func() {
				key := shamir3pass.GenerateKey(1024)
				holder, err := NewHolder(&key, Faces)
				Expect(err).To(BeNil())
				Expect(holder).ToNot(BeNil())
				err = holder.SetBothEncrypted(nil)
				Expect(err).ToNot(BeNil())
			})
		})
		Context("When the index is valid", func() {
			It("Should return an error", func() {
				key := shamir3pass.GenerateKey(1024)
				holder, err := NewHolder(&key, Faces)
				Expect(err).To(BeNil())
				Expect(holder).ToNot(BeNil())
				err = holder.SetBothEncrypted([]*big.Int{big.NewInt(0)})
				Expect(err).ToNot(BeNil())
			})
		})
	})
})
