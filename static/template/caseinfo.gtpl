						就诊日期：<div id="recordtime">{{.CreateTime}}</div><br>
						
						<input type="hidden" id="r_pid" name="r_pid" value="{{.PatientID}}"></input>
						<div class="pure-control-group">
							<label for="<MainComplaint">主&nbsp;&nbsp;&nbsp;&nbsp;诉</label>
							<textarea id="MainComplaint" name="MainComplaint" class="pure-input-2-3" rows="3" placeholder="主诉" disabled="disabled">{{.MainComplaint}}</textarea>
						</div>
						<div class="pure-control-group">
							<label for="ExamReport">检查报告</label>
							<textarea id="ExamReport" name="ExamReport" class="pure-input-2-3" rows="18" placeholder="检查报告" disabled="disabled">{{.ExamReport}}</textarea>
						</div>
						<div class="pure-control-group">
							<label for="Diag">诊&nbsp;&nbsp;&nbsp;&nbsp;断</label>
							<input type="text" id="Diag" name="Diag" class="pure-input-2-3" placeholder="诊断" value="{{.Diag}}" disabled="disabled"></input>
						</div>
						<div class="pure-control-group">
							<label for="DRR">医&nbsp;&nbsp;&nbsp;&nbsp;嘱</label>
							<textarea id="DRR" name="DRR" class="pure-input-2-3" rows="3" placeholder="医嘱" disabled="disabled">{{.DRR}}</textarea>
						</div>
						<div class="pure-control-group">
							<label for="Presciption">处&nbsp;&nbsp;&nbsp;&nbsp;方</label>
							<textarea id="Presciption" name="Presciption" class="pure-input-2-3" rows="3" placeholder="处方" disabled="disabled">{{.Presciption}}</textarea>
						</div>
						<div class="pure-control-group">
							<label for="Notes">备&nbsp;&nbsp;&nbsp;&nbsp;注</label>
							<textarea id="Notes" name="Notes" class="pure-input-2-3" rows="3" placeholder="备注" disabled="disabled">{{.Notes}}</textarea>
						</div>