package dynobuffers

import (
	"github.com/host6/dynobuffers"
	"math"
	flatbuffers "github.com/google/flatbuffers/go"
)

func fillArticleFlatBuffers(bl *flatbuffers.Builder) flatbuffers.UOffsetT {
	name := bl.CreateString("str")
	interName := bl.CreateString("str")
	pcText := bl.CreateString("str")
	fontName := bl.CreateString("str")
	omanText := bl.CreateString("str")
	pcBitmap := bl.CreateString("str")
	rmText := bl.CreateString("str")
	barcode := bl.CreateString("str")
	info := bl.CreateString("str")
	hhtFontName := bl.CreateString("str")
	aliasName := bl.CreateString("str")
	hqID := bl.CreateString("str")
	isActiveModifier := bl.CreateString("str")
	mlName := bl.CreateString("str")
	mlKSName := bl.CreateString("str")
	mlPCText := bl.CreateString("str")
	mlRMText := bl.CreateString("str")
	mlOmanText := bl.CreateString("str")

	ArticleStart(bl)
	ArticleAddId(bl, 1)
	ArticleAddArticleNumber(bl, 1)
	ArticleAddName(bl, name)
	ArticleAddInternalName(bl, interName)
	ArticleAddArticleManual(bl, true)
	ArticleAddArticleHash(bl, true)
	ArticleAddIdCourses(bl, 1)
	ArticleAddIdDepartament(bl, 1)
	ArticleAddPcBitmap(bl, pcBitmap)
	ArticleAddPcColor(bl, 1)
	ArticleAddPcText(bl, pcText)
	ArticleAddPcFontName(bl, fontName)
	ArticleAddPcFontSize(bl, 1)
	ArticleAddPcFontAttr(bl, 1)
	ArticleAddPcFontColor(bl, 1)
	ArticleAddRmText(bl, rmText)
	ArticleAddRmFontSize(bl, 1)
	ArticleAddIdPacking(bl, 1)
	ArticleAddIdCommission(bl, 1)
	ArticleAddIdPromotions(bl, 1)
	ArticleAddSavepoints(bl, 1)
	ArticleAddQuantity(bl, 1)
	ArticleAddHideonhold(bl, true)
	ArticleAddBarcode(bl, barcode)
	ArticleAddTimeActive(bl, true)
	ArticleAddAftermin(bl, 1)
	ArticleAddPeriodmin(bl, 1)
	ArticleAddRoundmin(bl, 1)
	ArticleAddIdCurrency(bl, 1)
	ArticleAddControlActive(bl, true)
	ArticleAddControlTime(bl, 1)
	ArticleAddPluNumberVanduijnen(bl, 1)
	ArticleAddSequence(bl, 1)
	ArticleAddRmSequence(bl, 1)
	ArticleAddPurchasePrice(bl, 1)
	ArticleAddIdVdGroup(bl, 1)
	ArticleAddMenu(bl, true)
	ArticleAddSensitive(bl, true)
	ArticleAddSensitiveOption(bl, true)
	ArticleAddDailyStock(bl, 1)
	ArticleAddInfo(bl, info)
	ArticleAddWarningLevel(bl, 1)
	ArticleAddFreeAfterPay(bl, 1)
	ArticleAddIdFoodGroup(bl, 1)
	ArticleAddArticleType(bl, true)
	ArticleAddIdInventoryItem(bl, 1)
	ArticleAddIdRecipe(bl, 1)
	ArticleAddIdUnitySales(bl, 1)
	ArticleAddCanSavepoints(bl, true)
	ArticleAddShowInKitchenScreen(bl, true)
	ArticleAddDecreaseSavepoints(bl, 1)
	ArticleAddHhtColor(bl, 1)
	ArticleAddHhtFontName(bl, hhtFontName)
	ArticleAddHhtFontSize(bl, 1)
	ArticleAddHhtFontAttr(bl, 1)
	ArticleAddHhtFontColor(bl, 1)
	ArticleAddTip(bl, true)
	ArticleAddIdBecoGroup(bl, 1)
	ArticleAddIdBecoLocation(bl, 1)
	ArticleAddBcStandardDosage(bl, 1)
	ArticleAddBcAlternativeDosage(bl, 1)
	ArticleAddBcDisablebalance(bl, true)
	ArticleAddBcUseLocations(bl, true)
	ArticleAddTimeRate(bl, 1)
	ArticleAddIdFreeOption(bl, 1)
	ArticleAddPartyArticle(bl, true)
	ArticleAddIdPuaGroups(bl, 1)
	ArticleAddPromo(bl, true)
	ArticleAddOneHandLimit(bl, 1)
	ArticleAddConsolidateQuantity(bl, 1)
	ArticleAddConsolidateAliasName(bl, aliasName)
	ArticleAddHqId(bl, hqID)
	ArticleAddIsActive(bl, true)
	ArticleAddIsActiveModified(bl, 1)
	ArticleAddIsActiveModifier(bl, isActiveModifier)
	ArticleAddRentPriceType(bl, true)
	ArticleAddIdRentalGroup(bl, 1)
	ArticleAddConditionCheckInOrder(bl, true)
	ArticleAddWeightRequired(bl, true)
	ArticleAddDailyNumeric1(bl, 1)
	ArticleAddDailyNumeric2(bl, 1)
	ArticleAddPrepMin(bl, 1)
	ArticleAddIdArticleKsp(bl, 1)
	ArticleAddWarnMin(bl, 1)
	ArticleAddEmptyArticle(bl, true)
	ArticleAddBcDebitcredit(bl, true)
	ArticleAddPrepSec(bl, 1)
	ArticleAddIdSuppliers(bl, 1)
	ArticleAddMainPrice(bl, true)
	ArticleAddOmanText(bl, omanText)
	ArticleAddIdAgeGroups(bl, 1)
	ArticleAddSurcharge(bl, true)
	ArticleAddInfoData(bl, 1)
	ArticleAddPosDisabled(bl, true)
	ArticleAddMlName(bl, mlName)
	ArticleAddMlKsName(bl, mlKSName)
	ArticleAddAltArticles(bl, 1)
	ArticleAddAltAlias(bl, 1)
	ArticleAddNeedPrep(bl, true)
	ArticleAddAutoOnhold(bl, true)
	ArticleAddIdKsWf(bl, 1)
	ArticleAddKsWfType(bl, 1)
	ArticleAddAskCourse(bl, true)
	ArticleAddPopupInfo(bl, 1)
	ArticleAddAllowOrderItems(bl, true)
	ArticleAddMustCombined(bl, true)
	ArticleAddBlockDiscount(bl, true)
	ArticleAddHasDefaultOptions(bl, true)
	ArticleAddHhtDefaultSetting(bl, true)
	ArticleAddOmanDefaultSetting(bl, true)
	ArticleAddIdRentPeriods(bl, 1)
	ArticleAddDelaySeparateMins(bl, 1)
	ArticleAddIdKsc(bl, 1)
	ArticleAddMlPcText(bl, mlPCText)
	ArticleAddMlRmText(bl, mlRMText)
	ArticleAddMlOmanText(bl, mlOmanText)
	ArticleAddPosArticleType(bl, true)
	ArticleAddSingleFreeOption(bl, true)
	ArticleAddKsSingleItem(bl, true)
	ArticleAddAllergen(bl, true)
	ArticleAddAutoResetcourse(bl, true)
	ArticleAddBlockTransfer(bl, true)
	ArticleAddIdSizeModifier(bl, 1)
	return ArticleEnd(bl)
}

func fillArticleDynoBuffer(bf *dynobuffers.Buffer) {
	bf.Set("id", int64(2))
	bf.Set("article_number", int32(1))
	bf.Set("name", "str")
	bf.Set("internal_name", "str")
	bf.Set("article_manual", true)
	bf.Set("article_hash", true)
	bf.Set("id_courses", int64(2))
	bf.Set("id_departament", int64(2))
	bf.Set("pc_bitmap", "str") // blob
	bf.Set("pc_color", int32(1))
	bf.Set("pc_text", "str")
	bf.Set("pc_font_name", "str")
	bf.Set("pc_font_size", int32(1))
	bf.Set("pc_font_attr", int32(1))
	bf.Set("pc_font_color", int32(1))
	bf.Set("rm_text", "str")
	bf.Set("rm_font_size", int32(1))
	bf.Set("id_packing", int64(2))
	bf.Set("id_commission", int64(2))
	bf.Set("id_promotions", int64(2))
	bf.Set("savepoints", int32(1))
	bf.Set("quantity", int32(1))
	bf.Set("hideonhold", true)
	bf.Set("barcode", "str")
	bf.Set("time_active", true)
	bf.Set("aftermin", int32(1))
	bf.Set("periodmin", int32(1))
	bf.Set("roundmin", int32(1))
	bf.Set("id_currency", int64(2))
	bf.Set("control_active", true)
	bf.Set("control_time", int32(1))
	bf.Set("plu_number_vanduijnen", int32(1))
	bf.Set("sequence", int32(1))
	bf.Set("rm_sequence", int32(1))
	bf.Set("purchase_price", float32(0.123))
	bf.Set("id_vd_group", int64(2))
	bf.Set("menu", true)
	bf.Set("sensitive", true)
	bf.Set("sensitive_option", true)
	bf.Set("daily_stock", int32(1))
	bf.Set("info", "str")
	bf.Set("warning_level", int32(1))
	bf.Set("free_after_pay", int32(1))
	bf.Set("id_food_group", int64(2))
	bf.Set("article_type", byte(3))
	bf.Set("id_inventory_item", int64(2))
	bf.Set("id_recipe", int64(2))
	bf.Set("id_unity_sales", int64(2))
	bf.Set("can_savepoints", true)
	bf.Set("show_in_kitchen_screen", true)
	bf.Set("decrease_savepoints", int32(1))
	bf.Set("hht_color", int32(1))
	bf.Set("hht_font_name", "str")
	bf.Set("hht_font_size", int32(1))
	bf.Set("hht_font_attr", int32(1))
	bf.Set("hht_font_color", int32(1))
	bf.Set("tip", true)
	bf.Set("id_beco_group", int64(2))
	bf.Set("id_beco_location", int64(2))
	bf.Set("bc_standard_dosage", int32(1))
	bf.Set("bc_alternative_dosage", int32(1))
	bf.Set("bc_disablebalance", true)
	bf.Set("bc_use_locations", true)
	bf.Set("time_rate", float32(0.123))
	bf.Set("id_free_option", int64(2))
	bf.Set("party_article", true)
	bf.Set("id_pua_groups", int64(2))
	bf.Set("promo", true)
	bf.Set("one_hand_limit", int32(1))
	bf.Set("consolidate_quantity", int32(1))
	bf.Set("consolidate_alias_name", "str")
	bf.Set("hq_id", "str")
	bf.Set("is_active", true)
	bf.Set("is_active_modified", int32(1)) // timestamp
	bf.Set("is_active_modifier", "str")
	bf.Set("rent_price_type", true)
	bf.Set("id_rental_group", int64(2))
	bf.Set("condition_check_in_order", true)
	bf.Set("weight_required", true)
	bf.Set("daily_numeric_1", float32(0.123))
	bf.Set("daily_numeric_2", float32(0.123))
	bf.Set("prep_min", int32(1))
	bf.Set("id_article_ksp", int64(2))
	bf.Set("warn_min", int32(1))
	bf.Set("empty_article", true)
	bf.Set("bc_debitcredit", true)
	bf.Set("prep_sec", int32(1))
	bf.Set("id_suppliers", int64(2))
	bf.Set("main_price", true)
	bf.Set("oman_text", "str")
	bf.Set("id_age_groups", int64(2))
	bf.Set("surcharge", true)
	bf.Set("info_data", "str") //blob
	bf.Set("pos_disabled", true)
	bf.Set("ml_name", "str")    // blob
	bf.Set("ml_ks_name", "str") // blob
	bf.Set("alt_articles", int64(2))
	bf.Set("alt_alias", "str")
	bf.Set("need_prep", true)
	bf.Set("auto_onhold", true)
	bf.Set("id_ks_wf", int64(2))
	bf.Set("ks_wf_type", int32(1))
	bf.Set("ask_course", true)
	bf.Set("popup_info", "str")
	bf.Set("allow_order_items", true)
	bf.Set("must_combined", true)
	bf.Set("block_discount", true)
	bf.Set("has_default_options", true)
	bf.Set("hht_default_setting", true)
	bf.Set("oman_default_setting", true)
	bf.Set("id_rent_periods", int64(2))
	bf.Set("delay_separate_mins", int32(1))
	bf.Set("id_ksc", int64(2))
	bf.Set("ml_pc_text", "str")   // blob
	bf.Set("ml_rm_text", "str")   // blob
	bf.Set("ml_oman_text", "str") // blob
	bf.Set("pos_article_type", true)
	bf.Set("single_free_option", true)
	bf.Set("ks_single_item", true)
	bf.Set("allergen", true)
	bf.Set("auto_resetcourse", true)
	bf.Set("block_transfer", true)
	bf.Set("id_size_modifier", int64(2))
}

func getArticleSchemaDynoBuffer() *dynobuffers.Schema {
	s := dynobuffers.NewSchema()
	s.AddField("id", dynobuffers.FieldTypeLong)
	s.AddField("article_number", dynobuffers.FieldTypeInt)
	s.AddField("name", dynobuffers.FieldTypeString)
	s.AddField("internal_name", dynobuffers.FieldTypeString)
	s.AddField("article_manual", dynobuffers.FieldTypeBool)
	s.AddField("article_hash", dynobuffers.FieldTypeBool)
	s.AddField("id_courses", dynobuffers.FieldTypeLong)
	s.AddField("id_departament", dynobuffers.FieldTypeLong)
	s.AddField("pc_bitmap", dynobuffers.FieldTypeString) // blob
	s.AddField("pc_color", dynobuffers.FieldTypeInt)
	s.AddField("pc_text", dynobuffers.FieldTypeString)
	s.AddField("pc_font_name", dynobuffers.FieldTypeString)
	s.AddField("pc_font_size", dynobuffers.FieldTypeInt)
	s.AddField("pc_font_attr", dynobuffers.FieldTypeInt)
	s.AddField("pc_font_color", dynobuffers.FieldTypeInt)
	s.AddField("rm_text", dynobuffers.FieldTypeString)
	s.AddField("rm_font_size", dynobuffers.FieldTypeInt)
	s.AddField("id_packing", dynobuffers.FieldTypeLong)
	s.AddField("id_commission", dynobuffers.FieldTypeLong)
	s.AddField("id_promotions", dynobuffers.FieldTypeLong)
	s.AddField("savepoints", dynobuffers.FieldTypeInt)
	s.AddField("quantity", dynobuffers.FieldTypeInt)
	s.AddField("hideonhold", dynobuffers.FieldTypeBool)
	s.AddField("barcode", dynobuffers.FieldTypeString)
	s.AddField("time_active", dynobuffers.FieldTypeBool)
	s.AddField("aftermin", dynobuffers.FieldTypeInt)
	s.AddField("periodmin", dynobuffers.FieldTypeInt)
	s.AddField("roundmin", dynobuffers.FieldTypeInt)
	s.AddField("id_currency", dynobuffers.FieldTypeLong)
	s.AddField("control_active", dynobuffers.FieldTypeBool)
	s.AddField("control_time", dynobuffers.FieldTypeInt)
	s.AddField("plu_number_vanduijnen", dynobuffers.FieldTypeInt)
	s.AddField("sequence", dynobuffers.FieldTypeInt)
	s.AddField("rm_sequence", dynobuffers.FieldTypeInt)
	s.AddField("purchase_price", dynobuffers.FieldTypeFloat)
	s.AddField("id_vd_group", dynobuffers.FieldTypeLong)
	s.AddField("menu", dynobuffers.FieldTypeBool)
	s.AddField("sensitive", dynobuffers.FieldTypeBool)
	s.AddField("sensitive_option", dynobuffers.FieldTypeBool)
	s.AddField("daily_stock", dynobuffers.FieldTypeInt)
	s.AddField("info", dynobuffers.FieldTypeString)
	s.AddField("warning_level", dynobuffers.FieldTypeInt)
	s.AddField("free_after_pay", dynobuffers.FieldTypeInt)
	s.AddField("id_food_group", dynobuffers.FieldTypeLong)
	s.AddField("article_type", dynobuffers.FieldTypeByte)
	s.AddField("id_inventory_item", dynobuffers.FieldTypeLong)
	s.AddField("id_recipe", dynobuffers.FieldTypeLong)
	s.AddField("id_unity_sales", dynobuffers.FieldTypeLong)
	s.AddField("can_savepoints", dynobuffers.FieldTypeBool)
	s.AddField("show_in_kitchen_screen", dynobuffers.FieldTypeBool)
	s.AddField("decrease_savepoints", dynobuffers.FieldTypeInt)
	s.AddField("hht_color", dynobuffers.FieldTypeInt)
	s.AddField("hht_font_name", dynobuffers.FieldTypeString)
	s.AddField("hht_font_size", dynobuffers.FieldTypeInt)
	s.AddField("hht_font_attr", dynobuffers.FieldTypeInt)
	s.AddField("hht_font_color", dynobuffers.FieldTypeInt)
	s.AddField("tip", dynobuffers.FieldTypeBool)
	s.AddField("id_beco_group", dynobuffers.FieldTypeLong)
	s.AddField("id_beco_location", dynobuffers.FieldTypeLong)
	s.AddField("bc_standard_dosage", dynobuffers.FieldTypeInt)
	s.AddField("bc_alternative_dosage", dynobuffers.FieldTypeInt)
	s.AddField("bc_disablebalance", dynobuffers.FieldTypeBool)
	s.AddField("bc_use_locations", dynobuffers.FieldTypeBool)
	s.AddField("time_rate", dynobuffers.FieldTypeFloat)
	s.AddField("id_free_option", dynobuffers.FieldTypeLong)
	s.AddField("party_article", dynobuffers.FieldTypeBool)
	s.AddField("id_pua_groups", dynobuffers.FieldTypeLong)
	s.AddField("promo", dynobuffers.FieldTypeBool)
	s.AddField("one_hand_limit", dynobuffers.FieldTypeInt)
	s.AddField("consolidate_quantity", dynobuffers.FieldTypeInt)
	s.AddField("consolidate_alias_name", dynobuffers.FieldTypeString)
	s.AddField("hq_id", dynobuffers.FieldTypeString)
	s.AddField("is_active", dynobuffers.FieldTypeBool)
	s.AddField("is_active_modified", dynobuffers.FieldTypeInt) // timestamp
	s.AddField("is_active_modifier", dynobuffers.FieldTypeString)
	s.AddField("rent_price_type", dynobuffers.FieldTypeBool)
	s.AddField("id_rental_group", dynobuffers.FieldTypeLong)
	s.AddField("condition_check_in_order", dynobuffers.FieldTypeBool)
	s.AddField("weight_required", dynobuffers.FieldTypeBool)
	s.AddField("daily_numeric_1", dynobuffers.FieldTypeFloat)
	s.AddField("daily_numeric_2", dynobuffers.FieldTypeFloat)
	s.AddField("prep_min", dynobuffers.FieldTypeInt)
	s.AddField("id_article_ksp", dynobuffers.FieldTypeLong)
	s.AddField("warn_min", dynobuffers.FieldTypeInt)
	s.AddField("empty_article", dynobuffers.FieldTypeBool)
	s.AddField("bc_debitcredit", dynobuffers.FieldTypeBool)
	s.AddField("prep_sec", dynobuffers.FieldTypeInt)
	s.AddField("id_suppliers", dynobuffers.FieldTypeLong)
	s.AddField("main_price", dynobuffers.FieldTypeBool)
	s.AddField("oman_text", dynobuffers.FieldTypeString)
	s.AddField("id_age_groups", dynobuffers.FieldTypeLong)
	s.AddField("surcharge", dynobuffers.FieldTypeBool)
	s.AddField("info_data", dynobuffers.FieldTypeString) //blob
	s.AddField("pos_disabled", dynobuffers.FieldTypeBool)
	s.AddField("ml_name", dynobuffers.FieldTypeString)    // blob
	s.AddField("ml_ks_name", dynobuffers.FieldTypeString) // blob
	s.AddField("alt_articles", dynobuffers.FieldTypeLong)
	s.AddField("alt_alias", dynobuffers.FieldTypeString)
	s.AddField("need_prep", dynobuffers.FieldTypeBool)
	s.AddField("auto_onhold", dynobuffers.FieldTypeBool)
	s.AddField("id_ks_wf", dynobuffers.FieldTypeLong)
	s.AddField("ks_wf_type", dynobuffers.FieldTypeInt)
	s.AddField("ask_course", dynobuffers.FieldTypeBool)
	s.AddField("popup_info", dynobuffers.FieldTypeString)
	s.AddField("allow_order_items", dynobuffers.FieldTypeBool)
	s.AddField("must_combined", dynobuffers.FieldTypeBool)
	s.AddField("block_discount", dynobuffers.FieldTypeBool)
	s.AddField("has_default_options", dynobuffers.FieldTypeBool)
	s.AddField("hht_default_setting", dynobuffers.FieldTypeBool)
	s.AddField("oman_default_setting", dynobuffers.FieldTypeBool)
	s.AddField("id_rent_periods", dynobuffers.FieldTypeLong)
	s.AddField("delay_separate_mins", dynobuffers.FieldTypeInt)
	s.AddField("id_ksc", dynobuffers.FieldTypeLong)
	s.AddField("ml_pc_text", dynobuffers.FieldTypeString)   // blob
	s.AddField("ml_rm_text", dynobuffers.FieldTypeString)   // blob
	s.AddField("ml_oman_text", dynobuffers.FieldTypeString) // blob
	s.AddField("pos_article_type", dynobuffers.FieldTypeBool)
	s.AddField("single_free_option", dynobuffers.FieldTypeBool)
	s.AddField("ks_single_item", dynobuffers.FieldTypeBool)
	s.AddField("allergen", dynobuffers.FieldTypeBool)
	s.AddField("auto_resetcourse", dynobuffers.FieldTypeBool)
	s.AddField("block_transfer", dynobuffers.FieldTypeBool)
	s.AddField("id_size_modifier", dynobuffers.FieldTypeLong)
	return s
}

func round(x float64, prec int) float64 {
	var rounder float64
	pow := math.Pow(10, float64(prec))
	intermed := x * pow
	_, frac := math.Modf(intermed)
	if frac >= 0.5 {
		rounder = math.Ceil(intermed)
	} else {
		rounder = math.Floor(intermed)
	}

	return rounder / pow
}
