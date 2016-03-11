package fbprototest

import (
	sample "fbproto/test"

	flatbuffers "github.com/google/flatbuffers/go"
)

func TestFBPrototest() {
	testMonster()
	testDvMonster()
}

func testMonster() {
	builder := flatbuffers.NewBuilder(0)

	weaponOne := builder.CreateString("Sword")
	weaponTwo := builder.CreateString("Axe")
	// Create the first `Weapon` ("Sword").
	sample.WeaponStart(builder)
	sample.WeaponAddName(builder, weaponOne)
	sample.WeaponAddDamage(builder, 3)
	sword := sample.WeaponEnd(builder)
	// Create the second `Weapon` ("Axe").
	sample.WeaponStart(builder)
	sample.WeaponAddName(builder, weaponTwo)
	sample.WeaponAddDamage(builder, 5)
	axe := sample.WeaponEnd(builder)

	// Serialize a name for our monster, called "Orc".
	name := builder.CreateString("Orc22")
	// Create a `vector` representing the inventory of the Orc. Each number
	// could correspond to an item that can be claimed after he is slain.
	// Note: Since we prepend the bytes, this loop iterates in reverse.
	sample.MonsterStartInventoryVector(builder, 10)
	for i := 9; i >= 0; i-- {
		builder.PrependByte(byte(i))
	}
	inv := builder.EndVector(10)

	// Create a FlatBuffer vector and prepend the weapons.
	// Note: Since we prepend the data, prepend them in reverse order.
	sample.MonsterStartWeaponsVector(builder, 2)
	builder.PrependUOffsetT(axe)
	builder.PrependUOffsetT(sword)
	weapons := builder.EndVector(2)

	// Create a `Vec3`, representing the Orc's position in 3-D space.
	pos := sample.CreateVec3(builder, 1.0, 2.0, 3.0)

	// Create our monster using `MonsterStart()` and `MonsterEnd()`.
	sample.MonsterStart(builder)
	sample.MonsterAddPos(builder, pos)
	sample.MonsterAddHp(builder, 408)
	sample.MonsterAddName(builder, name)
	sample.MonsterAddInventory(builder, inv)
	sample.MonsterAddColor(builder, sample.ColorRed)
	sample.MonsterAddWeapons(builder, weapons)
	sample.MonsterAddEquippedType(builder, sample.EquipmentWeapon)
	sample.MonsterAddEquipped(builder, axe)
	orc := sample.MonsterEnd(builder)

	// Call `Finish()` to instruct the builder that this monster is complete.
	builder.Finish(orc)

	// This must be called after `Finish()`.
	buf := builder.FinishedBytes() // Of type `byte[]`.
	println(len(buf))

	decodeMonster(buf)
}
func decodeMonster(buf []byte) {
	//--------------------------------------------------------------------
	monster := sample.GetRootAsMonster(buf, 0)
	hp := monster.Hp()
	mana := monster.Mana()
	name2 := string(monster.Name()) // Note: `monster.Name()` returns a byte[].
	println("======", hp, mana, name2)

	pos2 := monster.Pos(nil)
	x := int(pos2.X())
	y := int(pos2.Y())
	z := int(pos2.Z())
	println("==", x, y, z)
	// Note: Whenever you access a new object, like in `Pos()`, a new temporary
	// accessor object gets created. If your code is very performance sensitive,
	// you can pass in a pointer to an existing `Vec3` instead of `nil`. This
	// allows you to reuse it across many calls to reduce the amount of object
	// allocation/garbage collection.

	invLength := monster.InventoryLength()
	thirdItem := monster.Inventory(2)
	println("invLength=%d,thirdItem=%d", invLength, thirdItem)
	weaponLength := monster.WeaponsLength()
	weapon := new(sample.Weapon) // We need a `sample.Weapon` to pass into `monster.Weapons()`
	// to capture the output of the function.
	if monster.Weapons(weapon, 1) {
		secondWeaponName := string(weapon.Name())
		secondWeaponDamage := weapon.Damage()
		println(weaponLength, secondWeaponName, secondWeaponDamage)
	}

	// We need a `flatbuffers.Table` to capture the output of the
	// `monster.Equipped()` function.
	unionTable := new(flatbuffers.Table)
	if monster.Equipped(unionTable) {
		unionType := monster.EquippedType()
		if unionType == sample.EquipmentWeapon {
			// Create a `sample.Weapon` object that can be initialized with the contents
			// of the `flatbuffers.Table` (`unionTable`), which was populated by
			// `monster.Equipped()`.
			unionWeapon := new(sample.Weapon)
			unionWeapon.Init(unionTable.Bytes, unionTable.Pos)
			weaponName := string(unionWeapon.Name())
			weaponDamage := unionWeapon.Damage()

			println(weaponName)
			println(weaponDamage)
		}
	}
}

func testDvMonster() {
	builder := flatbuffers.NewBuilder(0)

	name := builder.CreateString("davin")

	sample.DvMonsterStart(builder)
	sample.DvMonsterAddId(builder, 1200)
	sample.DvMonsterAddName(builder, name)

	dvTest := sample.DvMonsterEnd(builder)
	builder.Finish(dvTest)

	buf := builder.FinishedBytes()
	println("长度=", len(buf))

	dvMonster := sample.GetRootAsDvMonster(buf, 0)
	dvName := string(dvMonster.Name())
	dvId := dvMonster.Id()
	println(dvName, dvId)
}
